package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"math/big"

	revm_api "github.com/0xEyrie/revmffi/core"
	"github.com/ethereum/go-ethereum/core"

	"github.com/0xEyrie/revmffi/core/contracts/erc20"
	"github.com/0xEyrie/revmffi/core/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	ethState "github.com/ethereum/go-ethereum/core/state"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/triedb"
)

const cancun = 17
const caller = "0x10"
const gaslimit = 3000000

func main() {
	erc20abi, _ := erc20.Erc20MetaData.GetAbi()
	erc20bin, _ := hexutil.Decode(erc20.Erc20Bin)
	callerAddr := common.HexToAddress(caller)
	recipientAddr := common.HexToAddress("0x20")
	// Create VM
	memdb := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(memdb, triedb.HashDefaults)
	kvstore := ethState.NewDatabase(tdb, nil)

	stateDB, _ := state.New(ethTypes.EmptyRootHash, kvstore)
	acnt := ethTypes.NewEmptyStateAccount()
	acnt.Balance = uint256.NewInt(1000000000000000000)
	stateDB.SetAccount(callerAddr, acnt)
	stateDB.SetAccount(recipientAddr, acnt)
	stateDB.SetAccount(common.Address{}, acnt)
	random := common.Hash{1} // 设置一个非零的随机值
	evm := revm_api.NewEVM(vm.BlockContext{
		BlockNumber: big.NewInt(0),
		Random:      &random,
		GasLimit:    gaslimit,
		BlobBaseFee: big.NewInt(0),
		//Coinbase:    common.HexToAddress(caller),
	}, *stateDB, revm_api.Config{
		Spec:      16,
		NoBaseFee: true, // 禁用基础费用检查
	})
	defer evm.Destroy()
	// ERC20 create
	packedData, _ := erc20abi.Constructor.Inputs.Pack("Mock", "Mock")
	calldata := append(erc20bin, packedData...)
	//txcontext := testutils.MockTx(callerAddr, common.Address{}, calldata, 0)
	//block := testutils.MockBlock(1)
	result, err := evm.Execute(vm.AccountRef(callerAddr), &core.Message{
		From:          callerAddr,
		To:            nil,
		Nonce:         0,
		GasLimit:      gaslimit,
		GasPrice:      big.NewInt(10000),
		GasFeeCap:     big.NewInt(0),
		GasTipCap:     big.NewInt(0),
		Data:          calldata,
		AccessList:    ethTypes.AccessList{},
		BlobHashes:    nil,
		BlobGasFeeCap: nil,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.String())
	//fmt.Println(result.GetResult())
	erc20Addr := common.Address(result.GetSuccess().Output.GetCreate().GetCreatedAddress())
	fmt.Println(erc20Addr.Hex())
	// ERC20 Mint
	mintData, _ := erc20abi.Pack("mint", callerAddr, big.NewInt(1000))
	result, err = evm.Execute(vm.AccountRef(callerAddr), &core.Message{
		From:          callerAddr,
		To:            &erc20Addr,
		Nonce:         1,
		GasLimit:      gaslimit,
		GasPrice:      big.NewInt(10000),
		GasFeeCap:     big.NewInt(0),
		GasTipCap:     big.NewInt(0),
		Data:          mintData,
		AccessList:    ethTypes.AccessList{},
		BlobHashes:    nil,
		BlobGasFeeCap: nil,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.String())
	fmt.Println("mint: ", result.GetResult())
	fmt.Println(stateDB.GetStateObject(callerAddr))

	// ERC20 Transfer

	transferData, _ := erc20abi.Pack("transfer", recipientAddr, big.NewInt(100))
	result, err = evm.Execute(vm.AccountRef(callerAddr), &core.Message{
		From: callerAddr,
		To:   &recipientAddr,

		Nonce:         2,
		GasLimit:      gaslimit,
		GasPrice:      big.NewInt(10000),
		GasFeeCap:     big.NewInt(0),
		GasTipCap:     big.NewInt(0),
		Data:          transferData,
		AccessList:    ethTypes.AccessList{},
		BlobHashes:    nil,
		BlobGasFeeCap: nil,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.String())
	fmt.Println("transfer: ", result.GetResult())
	fmt.Println(stateDB.GetStateObject(callerAddr))
	fmt.Println(stateDB.GetStateObject(recipientAddr))

	// ERC20 BalanceOf
	balanceOfData, _ := erc20abi.Pack("balanceOf", recipientAddr)
	result, err = evm.Execute(vm.AccountRef(callerAddr), &core.Message{
		From:          recipientAddr,
		To:            &erc20Addr,
		Nonce:         3,
		GasLimit:      gaslimit,
		GasPrice:      big.NewInt(10000),
		GasFeeCap:     big.NewInt(0),
		GasTipCap:     big.NewInt(0),
		Data:          balanceOfData,
		AccessList:    ethTypes.AccessList{},
		BlobHashes:    nil,
		BlobGasFeeCap: nil,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.String())
	fmt.Println("balanceOf: ", result.GetResult())
	fmt.Println(result.GetSuccess().Output.String())
	//fmt.Println(stateDB.GetStorage(callerAddr))
	fmt.Println(stateDB.GetStateObject(recipientAddr))

}
