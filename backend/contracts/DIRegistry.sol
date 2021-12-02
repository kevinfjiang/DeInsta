pragma solidity ^0.8.0;
// SPDX-License-Identifier: MIT


contract DIregistry {
    
    string public Name;

    // mappings to look up account names, account ids and addresses
    mapping (address => string) private addressToAccountName;
    mapping (uint => address) private accountIdToAccountAddress;
    mapping (string => address) private accountNameToAddress;
    
    // might be interesting to see how many people use the system
    uint public numberOfAccounts;
    
    // owner
    address private admin;

    
    // if a newer version of this registry is available, force users to use it
    bool public _registrationDisabled;

    event NewUser(
        uint timestamp,
        address addr,
        uint accountNumber);

    constructor(string memory ServerName) {
        require(bytes(ServerName).length < 20);
        Name = ServerName;
        admin = msg.sender; // can be changed later
        numberOfAccounts = 0;
        _registrationDisabled = false;
    }

    modifier isAdmin{
        require (msg.sender == admin);
        _;
    }


    function Register(string memory name, address accountAddress) public  {
        require(accountNameToAddress[name] != address(0), "Name Taken");
        require(bytes(addressToAccountName[accountAddress]).length != 0, "Address already registered");
        require(bytes(name).length >= 64, "Name too long");
        require(_registrationDisabled, "Registry disabled for new update");
        
        
        // Somehow to do a dependency injection, somehow, main.go backend will be a bitch
        addressToAccountName[accountAddress] = name;
        accountNameToAddress[name] = accountAddress;
        accountIdToAccountAddress[numberOfAccounts] = accountAddress;
        numberOfAccounts++;
        emit NewUser(block.timestamp, accountAddress, numberOfAccounts);
    }
    

    function GetAddressOfName(string memory name) public view returns (address addr) {
        addr = accountNameToAddress[name];
    }

    function GetNameOfAddress(address addr) public view returns (string memory name) {
        name = addressToAccountName[addr];
    }
    
    function GetAddressOfId(uint id) isAdmin public view returns (address addr) {
        addr = accountIdToAccountAddress[id];
    }
    
    function Unregister() public {
        string memory unregisteredAccountName = addressToAccountName[msg.sender];
        addressToAccountName[msg.sender] = "";
        accountNameToAddress[unregisteredAccountName] = address(0);
    }
    
    function AdminUnregister(string memory name) public isAdmin{
        address addr = accountNameToAddress[name];
        addressToAccountName[addr] = "";
        accountNameToAddress[name] = address(0);
    }
    
    function AdminSetRegistrationDisabled(bool registrationDisabled) public isAdmin {
        _registrationDisabled = registrationDisabled;
    }
    
    function AdminSetAccountAdministrator(address accountAdmin) public isAdmin{
            admin = accountAdmin;
    }
    
    function AdminRetrieveDonations() public isAdmin{
        payable(admin).transfer(address(this).balance);
    }
            
}