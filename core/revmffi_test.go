package core_test

import (
	"github.com/0xEyrie/revmffi/core/state"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/rethmint/revm-api/testutils"
	"math/big"
	"testing"

	"github.com/0xEyrie/revmffi/core/contracts/erc20"
	"github.com/0xEyrie/revmffi/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethState "github.com/ethereum/go-ethereum/core/state"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"
)

const cancun = 17
const caller = "0x10"

func Test_ERC20_Benchmark(t *testing.T) {
	erc20abi, _ := erc20.Erc20MetaData.GetAbi()
	erc20bin, _ := hexutil.Decode(erc20.Erc20Bin)
	callerAddr := common.HexToAddress(caller)
	// Create VM
	memdb := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(memdb, triedb.HashDefaults)
	kvstore := ethState.NewDatabase(tdb, nil)

	stateDB, _ := state.New(ethTypes.EmptyRootHash, kvstore)
	evm := NewEVM(vm.BlockContext{}, stateDB, Config{})
	defer evm.Destroy()
	// ERC20 create
	packedData, _ := erc20abi.Constructor.Inputs.Pack("Mock", "Mock")
	calldata := append(erc20bin, packedData...)
	txcontext := testutils.MockTx(callerAddr, common.Address{}, calldata, 0)
	block := testutils.MockBlock(1)
	result, err := evm.Execute(block.ToSerialized(), txcontext.ToSerialized())
	require.NoError(t, err)
	res, err := result.ProcessExecutionResult()
	require.NoError(t, err)
	success, ok := res.(types.Success)
	require.True(t, ok)
	erc20Addr := success.Output.DeployedAddress
	// ERC20 Mint
	mintData, _ := erc20abi.Pack("mint", callerAddr, big.NewInt(1000))
	txcontext = testutils.MockTx(callerAddr, erc20Addr, mintData, 1)
	result, err = evm.ExecuteTx(kvstore, block.ToSerialized(), txcontext.ToSerialized())
	require.NoError(t, err)
	res, err = result.ProcessExecutionResult()
	require.NoError(t, err)
	_, ok = res.(types.Success)
	require.True(t, ok)

	// ERC20 Transfer
	recipientAddr := common.HexToAddress("0x20")
	transferData, _ := erc20abi.Pack("transfer", recipientAddr, big.NewInt(100))
	txcontext = testutils.MockTx(callerAddr, erc20Addr, transferData, 2)
	result, err = evm.ExecuteTx(kvstore, block.ToSerialized(), txcontext.ToSerialized())
	require.NoError(t, err)
	res, err = result.ProcessExecutionResult()
	require.NoError(t, err)
	_, ok = res.(types.Success)
	require.True(t, ok)

	// ERC20 BalanceOf
	balanceOfData, _ := erc20abi.Pack("balanceOf", recipientAddr)
	txcontext = testutils.MockTx(callerAddr, erc20Addr, balanceOfData, 3)
	result, _ = evm.QueryTx(kvstore, block.ToSerialized(), txcontext.ToSerialized())
	res, err = result.ProcessExecutionResult()
	require.NoError(t, err)
	queryRes, ok := res.(types.Success)
	require.True(t, ok)
	balance := new(big.Int).SetBytes(queryRes.Output.Output).Uint64()
	require.Equal(t, uint64(100), balance)

}
