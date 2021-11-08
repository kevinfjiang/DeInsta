
contract ERRegistry {
    
    string public Name;
    string public Description;

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

    constructor(string memory SubName, string memory SubDescription) {
        require(bytes(SubName).length < 20);
        Name = SubName;
        Description = SubDescription;
        admin = msg.sender; // can be changed later
        numberOfAccounts = 0;
        _registrationDisabled = false;
    }

    modifier isAdmin{
        require (msg.sender == admin);
        _;
    }


    function register(string memory name, address accountAddress) public  {
        require(accountNameToAddress[name] != address(0), "Name Taken");
        require(bytes(addressToAccountName[accountAddress]).length != 0, "Address already registered");
        require(bytes(name).length >= 64, "Name too long");
        require(_registrationDisabled, "Registry disabled for new update");
        
        
        // Somehow to do a dependency injection, somehow, main.go backend will be a bitch
        addressToAccountName[accountAddress] = name;
        accountNameToAddress[name] = accountAddress;
        accountIdToAccountAddress[numberOfAccounts] = accountAddress;
        numberOfAccounts++;
    }
    

    function getAddressOfName(string memory name) public view returns (address addr) {
        addr = accountNameToAddress[name];
    }

    function getNameOfAddress(address addr) public view returns (string memory name) {
        name = addressToAccountName[addr];
    }
    
    function getAddressOfId(uint id) isAdmin public view returns (address addr) {
        addr = accountIdToAccountAddress[id];
    }
    
    function unregister() public {
        string memory unregisteredAccountName = addressToAccountName[msg.sender];
        addressToAccountName[msg.sender] = "";
        accountNameToAddress[unregisteredAccountName] = address(0);
    }
    
    function adminUnregister(string memory name) public isAdmin{
        address addr = accountNameToAddress[name];
        addressToAccountName[addr] = "";
        accountNameToAddress[name] = address(0);
    }
    
    function adminSetRegistrationDisabled(bool registrationDisabled) public isAdmin {
        _registrationDisabled = registrationDisabled;
    }
    
    function adminSetAccountAdministrator(address accountAdmin) public isAdmin{
            admin = accountAdmin;
    }
    
    function adminRetrieveDonations() public isAdmin{
        payable(admin).transfer(address(this).balance);
    }
            
}