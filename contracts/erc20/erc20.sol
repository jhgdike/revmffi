// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Erc20 {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    address public owner;
    
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Mint(address indexed to, uint256 value);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    constructor(string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
        decimals = 18;
        owner = msg.sender;
        totalSupply = 1000000 * 10**uint256(decimals);
        balanceOf[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balanceOf[msg.sender] >= _value, "Insufficient balance");
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(balanceOf[_from] >= _value, "Insufficient balance");
        require(allowance[_from][msg.sender] >= _value, "Insufficient allowance");
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        allowance[_from][msg.sender] -= _value;
        emit Transfer(_from, _to, _value);
        return true;
    }

    // 新增 mint 功能
    function mint(address _to, uint256 _value) public onlyOwner returns (bool success) {
        require(_to != address(0), "Invalid address");
        totalSupply += _value;
        balanceOf[_to] += _value;
        emit Mint(_to, _value);
        emit Transfer(address(0), _to, _value);
        return true;
    }

    // 批量铸造代币
    function mintBatch(address[] calldata _to, uint256[] calldata _value) public onlyOwner returns (bool success) {
        require(_to.length == _value.length, "Array lengths must match");
        for(uint i = 0; i < _to.length; i++) {
            require(_to[i] != address(0), "Invalid address");
            totalSupply += _value[i];
            balanceOf[_to[i]] += _value[i];
            emit Mint(_to[i], _value[i]);
            emit Transfer(address(0), _to[i], _value[i]);
        }
        return true;
    }

    // 查询多个地址的余额
    function balanceOfBatch(address[] calldata _accounts) public view returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](_accounts.length);
        for(uint i = 0; i < _accounts.length; i++) {
            balances[i] = balanceOf[_accounts[i]];
        }
        return balances;
    }

    // 转移所有权
    function transferOwnership(address _newOwner) public onlyOwner {
        require(_newOwner != address(0), "Invalid new owner");
        owner = _newOwner;
    }
} 