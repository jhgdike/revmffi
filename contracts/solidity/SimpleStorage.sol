// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    struct Data {
        string name;
        uint256 value;
        bool active;
        uint256 timestamp;
    }

    mapping(address => Data) private userData;
    address[] private users;

    event DataStored(address indexed user, string name, uint256 value, bool active);
    event DataUpdated(address indexed user, string name, uint256 value, bool active);

    function storeData(string memory _name, uint256 _value, bool _active) public {
        Data storage data = userData[msg.sender];
        data.name = _name;
        data.value = _value;
        data.active = _active;
        data.timestamp = block.timestamp;

        if (data.timestamp == block.timestamp) {
            users.push(msg.sender);
        }

        emit DataStored(msg.sender, _name, _value, _active);
    }

    function updateData(string memory _name, uint256 _value, bool _active) public {
        require(userData[msg.sender].timestamp > 0, "No data exists for this user");
        
        Data storage data = userData[msg.sender];
        data.name = _name;
        data.value = _value;
        data.active = _active;
        data.timestamp = block.timestamp;

        emit DataUpdated(msg.sender, _name, _value, _active);
    }

    function getData() public view returns (
        string memory name,
        uint256 value,
        bool active,
        uint256 timestamp
    ) {
        Data storage data = userData[msg.sender];
        return (data.name, data.value, data.active, data.timestamp);
    }

    function getUserCount() public view returns (uint256) {
        return users.length;
    }

    function getUserAtIndex(uint256 _index) public view returns (address) {
        require(_index < users.length, "Index out of bounds");
        return users[_index];
    }
} 