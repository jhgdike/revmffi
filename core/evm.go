package core

import (
	"math/big"

	"github.com/0xEyrie/revmffi/core/state"
	revmtypes "github.com/0xEyrie/revmffi/core/types"
	revm "github.com/0xEyrie/revmffi/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"google.golang.org/protobuf/proto"
)

type Config struct {
	Spec revm.SpecId

	NoBaseFee         bool
	hasCompiler       bool
	thershold         uint64
	maxConcurrentSize uint
}

// VM struct is the core of initiavm.
type EVM struct {
	Inner   revm.EVM
	Context vm.BlockContext
	// virtual machine configuration options used to initialise the evm.
	Config Config
	vm.TxContext
}

// NewVM return VM instance
func NewEVM(blockCtx vm.BlockContext, statedb state.ExtendedStateDB, config Config) EVM {
	var inner revm.EVM
	if config.hasCompiler {
		inner = revm.NewEVMWithCompiler(statedb, config.thershold, config.maxConcurrentSize, config.Spec)
	} else {
		inner = revm.NewEVM(statedb, config.Spec)
	}

	return EVM{
		Inner:   inner,
		Context: blockCtx,
		Config:  config,
	}
}

func (evm *EVM) GetSpecId() revm.SpecId {
	return evm.Config.Spec
}

func (evm *EVM) SetTxContext(txCtx vm.TxContext) {
	evm.TxContext = txCtx
}

func (evm *EVM) Destroy() {
	revm.DestroyVM(evm.Inner)
}

func (evm *EVM) SetBlockHashFn(hashFn vm.GetHashFunc) {
	evm.Inner.StateDB.SetBlockHashFn(hashFn)
}

func (evm *EVM) SetBlockNumber(number uint64) {
	evm.Inner.StateDB.SetBlockNumber(number)
}

// safeBigBytes returns the bytes of a big.Int, or zero if nil
func safeBigBytes(b *big.Int) []byte {
	if b == nil {
		return big.NewInt(0).Bytes()
	}
	return b.Bytes()
}

// Call execute transaction based on revm
// this function only support entry call of transactions
func (evm *EVM) Execute(
	caller vm.ContractRef, msg *core.Message,
) (*revmtypes.EvmResult, error) {
	// save block context on evm
	//var excessBlobGas uint64
	//if evm.Context.BlobBaseFee != nil {
	//	excessBlobGas = evm.Context.BlobBaseFee.Uint64()
	//}
	number := evm.Context.BlockNumber
	evm.SetBlockNumber(number.Uint64())
	block := &revmtypes.Block{
		Number:     safeBigBytes(number),
		Coinbase:   evm.Context.Coinbase.Bytes(),
		Timestamp:  safeBigBytes(big.NewInt(int64(evm.Context.Time))),
		GasLimit:   safeBigBytes(big.NewInt(int64(evm.Context.GasLimit))),
		Basefee:    safeBigBytes(evm.Context.BaseFee),
		Difficulty: safeBigBytes(evm.Context.Difficulty),
		Prevrandao: evm.Context.Random[:],
		//ExcessBlobGas: &excessBlobGas,
	}

	blockBuf, err := proto.Marshal(block)
	if err != nil {
		return nil, err
	}
	transaction := revmtypes.Transaction{
		Caller:         caller.Address().Bytes(),
		GasLimit:       msg.GasLimit,
		GasPrice:       safeBigBytes(msg.GasPrice),
		Nonce:          nil,
		TransactTo:     make([]byte, 20),
		Value:          make([]byte, 0),
		Data:           msg.Data,
		GasPriorityFee: safeBigBytes(msg.GasTipCap),
		AccessList: func(accl types.AccessList) []*revmtypes.AccessListItem {
			result := make([]*revmtypes.AccessListItem, len(accl))
			for i, acc := range accl {
				storageKeys := make([]*revmtypes.StorageKey, len(acc.StorageKeys))
				for j, key := range acc.StorageKeys {
					storageKeys[j] = &revmtypes.StorageKey{Value: key.Bytes()}
				}
				result[i] = &revmtypes.AccessListItem{
					Address:     acc.Address.Bytes(),
					StorageKeys: storageKeys,
				}
			}
			return result
		}(msg.AccessList),
		BlobHashes: func(hashes []common.Hash) [][]byte {
			result := make([][]byte, len(hashes))
			for i, hash := range hashes {
				result[i] = hash.Bytes()
			}
			return result
		}(msg.BlobHashes),
		MaxFeePerBlobGas:  nil,
		AuthorizationList: nil,
	}
	if msg.BlobGasFeeCap != nil {
		transaction.MaxFeePerBlobGas = safeBigBytes(msg.BlobGasFeeCap)
	}
	if msg.To != nil {
		transaction.TransactTo = msg.To.Bytes()
	}
	if msg.Value != nil {
		transaction.Value = msg.Value.Bytes()
	}

	txBuf, err := proto.Marshal(&transaction)
	if err != nil {
		return nil, err
	}

	res, err := evm.Inner.Execute(
		&blockBuf,
		&txBuf,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
