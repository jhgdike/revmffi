package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/rethmint/revm-api/erc20"
)

const (
	caller    = "0x10"
	gaslimit  = 3000000
	recipient = "0x20"
)

func main() {
	// 1. 准备合约 ABI 和字节码
	erc20abi, _ := erc20.Erc20MetaData.GetAbi()
	erc20bin, _ := hexutil.Decode(erc20.Erc20Bin)
	callerAddr := common.HexToAddress(caller)
	recipientAddr := common.HexToAddress(recipient)

	// 2. 创建 EVM 环境
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

	// 3. 创建交易消息
	msg := &core.Message{
		From:       callerAddr,
		To:         &common.Address{},
		Nonce:      0,
		GasLimit:   gaslimit,
		GasPrice:   big.NewInt(10000),
		GasFeeCap:  big.NewInt(0),
		GasTipCap:  big.NewInt(0),
		Data:       erc20bin,
		AccessList: evmtypes.AccessList{},
	}

	// 4. 创建状态数据库
	memDb := rawdb.NewMemoryDatabase()
	trieDb := triedb.NewDatabase(memDb, nil)
	statedb, _ := revmState.New(common.Hash{}, state.NewDatabase(trieDb, nil))

	// 5. 创建 EVM 实例
	evm := core.NewEVM(blockContext, *statedb, core.Config{spec: revmVm.CANCUN})

	// 6. 部署 ERC20 合约
	packedData, err := erc20abi.Constructor.Inputs.Pack("Mock Token", "Mock")
	if err != nil {
		fmt.Printf("打包构造函数参数失败: %v\n", err)
		return
	}
	calldata := append(erc20bin, packedData...)
	msg.Data = calldata

	// 7. 执行合约部署
	result, err := evm.Execute(vm.AccountRef(callerAddr), msg)
	if err != nil {
		fmt.Printf("部署合约失败: %v\n", err)
		return
	}
	contractAddress := result.String()
	fmt.Printf("合约地址: %s\n", contractAddress)

	// 8. 调用 mint 函数
	mintAmount := big.NewInt(1000)
	mintData, err := erc20abi.Pack("mint", callerAddr, mintAmount)
	if err != nil {
		fmt.Printf("打包 mint 函数参数失败: %v\n", err)
		return
	}
	msg.Data = mintData
	contractAddr := common.HexToAddress(contractAddress)
	msg.To = &contractAddr

	result, err = evm.Execute(vm.AccountRef(callerAddr), msg)
	if err != nil {
		fmt.Printf("调用 mint 失败: %v\n", err)
		return
	}
	fmt.Printf("Mint 结果: %s\n", result.String())

	// 9. 查询余额
	balanceData, err := erc20abi.Pack("balanceOf", callerAddr)
	if err != nil {
		fmt.Printf("打包 balanceOf 函数参数失败: %v\n", err)
		return
	}
	msg.Data = balanceData

	result, err = evm.Execute(vm.AccountRef(callerAddr), msg)
	if err != nil {
		fmt.Printf("查询余额失败: %v\n", err)
		return
	}
	fmt.Printf("余额: %s\n", result.String())

	// 10. 转账
	transferAmount := big.NewInt(100)
	transferData, err := erc20abi.Pack("transfer", recipientAddr, transferAmount)
	if err != nil {
		fmt.Printf("打包 transfer 函数参数失败: %v\n", err)
		return
	}
	msg.Data = transferData

	result, err = evm.Execute(vm.AccountRef(callerAddr), msg)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
		return
	}
	fmt.Printf("转账结果: %s\n", result.String())

	// 11. 查询接收者余额
	balanceData, err = erc20abi.Pack("balanceOf", recipientAddr)
	if err != nil {
		fmt.Printf("打包 balanceOf 函数参数失败: %v\n", err)
		return
	}
	msg.Data = balanceData

	result, err = evm.Execute(vm.AccountRef(callerAddr), msg)
	if err != nil {
		fmt.Printf("查询接收者余额失败: %v\n", err)
		return
	}
	fmt.Printf("接收者余额: %s\n", result.String())
}
