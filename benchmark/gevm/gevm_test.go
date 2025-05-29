package gevm_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
	"github.com/rethmint/revm-api/benchmark/contracts/erc20"
	"github.com/stretchr/testify/require"
)

const caller = "0x10"
const gaslimit = 3000000

func Test_ERC20_Benchmark(t *testing.T) {
	erc20abi, _ := erc20.Erc20MetaData.GetAbi()
	erc20bin, _ := hexutil.Decode(erc20.Erc20Bin)
	callerAddr := common.HexToAddress(caller)
	// Create VM
	config := params.MainnetChainConfig
	genesis := &core.Genesis{
		Config:     config,
		Coinbase:   common.Address{},
		Difficulty: big.NewInt(0),
		GasLimit:   gaslimit,
		Number:     config.LondonBlock.Uint64(),
		Timestamp:  *config.CancunTime,
		Alloc:      evmtypes.GenesisAlloc{},
	}
	blockContext := core.NewEVMBlockContext(genesis.ToBlock().Header(), nil, &common.Address{})
	txContext := core.NewEVMTxContext(&core.Message{
		From:       common.HexToAddress(caller),
		To:         &common.Address{},
		Nonce:      0,
		GasLimit:   gaslimit,
		GasPrice:   big.NewInt(10000),
		GasFeeCap:  big.NewInt(0),
		GasTipCap:  big.NewInt(0),
		Data:       erc20bin,
		AccessList: evmtypes.AccessList{},
	})

	memDb := rawdb.NewMemoryDatabase()
	trieDb := triedb.NewDatabase(memDb, nil)
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(trieDb, nil))
	evm := vm.NewEVM(blockContext, txContext, statedb, config, vm.Config{})

	// ERC20 create
	packedData, err := erc20abi.Constructor.Inputs.Pack("Mock Token", "Mock")
	require.NoError(t, err)
	calldata := append(erc20bin, packedData...)
	_, contractAddress, _, err := evm.Create(vm.AccountRef(callerAddr), calldata, gaslimit, new(uint256.Int))
	require.NoError(t, err)

	// 检查 owner 是否正确设置
	ownerData, _ := erc20abi.Pack("owner")
	ret, _, err := evm.StaticCall(vm.AccountRef(callerAddr), contractAddress, ownerData, gaslimit)
	require.NoError(t, err)
	require.Equal(t, callerAddr.Bytes(), ret[len(ret)-20:])

	// 由于调用者是合约创建者，所以也是 owner，可以直接调用 mint
	mintAmount := big.NewInt(1000)
	mintData, _ := erc20abi.Pack("mint", callerAddr, mintAmount)
	_, leftOverGas, err := evm.Call(vm.AccountRef(callerAddr), contractAddress, mintData, gaslimit, new(uint256.Int))
	require.NoError(t, err)
	fmt.Println("mint:", gaslimit-leftOverGas)

	// 检查余额是否正确增加
	balanceData, _ := erc20abi.Pack("balanceOf", callerAddr)
	ret, leftOverGas, err = evm.StaticCall(vm.AccountRef(callerAddr), contractAddress, balanceData, gaslimit)
	require.NoError(t, err)
	fmt.Println("balanceOf:", gaslimit-leftOverGas)
	balance := new(big.Int).SetBytes(ret)
	initialSupply := new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18)) // 1e24
	expectedBalance := new(big.Int).Add(initialSupply, mintAmount)
	require.Equal(t, expectedBalance, balance)

	// ERC20 Transfer
	recipientAddr := common.HexToAddress("0x20")
	transferAmount := big.NewInt(100)
	transferData, _ := erc20abi.Pack("transfer", recipientAddr, transferAmount)
	_, leftOverGas, err = evm.Call(vm.AccountRef(callerAddr), contractAddress, transferData, gaslimit, new(uint256.Int))
	require.NoError(t, err)
	fmt.Println("transfer:", gaslimit-leftOverGas)

	// 检查接收者余额
	balanceData, _ = erc20abi.Pack("balanceOf", recipientAddr)
	ret, leftOverGas, err = evm.StaticCall(vm.AccountRef(callerAddr), contractAddress, balanceData, gaslimit)
	require.NoError(t, err)
	balance = new(big.Int).SetBytes(ret)
	require.Equal(t, transferAmount, balance)
	fmt.Println("balanceOf:", gaslimit-leftOverGas)

	// 打印调试信息
	//fmt.Printf("合约地址: %s\n", contractAddress.Hex())
	//fmt.Printf("调用者地址: %s\n", callerAddr.Hex())
	//fmt.Printf("接收者地址: %s\n", recipientAddr.Hex())
}
