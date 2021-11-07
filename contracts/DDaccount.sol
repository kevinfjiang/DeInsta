pragma solidity ^0.8.0;
// SPDX-License-Identifier: MIT

contract DDAccount {
    string public Name;

    address[] public Posts;
    int16 public Karma; // Add more robust Karma system
    
    // total posts
    uint public CountPosts;
    
    address owner;
    
    constructor(string memory Name_) {
        Name = Name_;
        CountPosts = 0;
        owner = tx.origin;
    }
    
    modifier isAdmin(){
        require(msg.sender == owner);
        _;
    }
    
    // create new tweet
    function PostMessage(string memory postString) public isAdmin {
        require(bytes(postString).length < 500, "Post exceeds character limit"); 
        Post TempPost = new Post(postString, CountPosts++); // Dependency injectiooon or no?
        Posts.push(address(TempPost));
        
    }

    function DeletePost(uint PostID) external isAdmin{
        Post(payable(Posts[PostID])).DeletePost();
        delete Posts[PostID];
    }

    function GetPost(uint PostID) public view returns(address Post_){
        // Rewrite verbose returns of every item
        return Posts[PostID];
    }
    
    function GetOwnerAddress() public view returns (address adminAddress) {
        return owner;
    }
    
    function adminDeleteAccount() public payable isAdmin {
        selfdestruct(payable(owner));
    }    
}

contract Post{
    uint private timestamp;
    uint public PostNumber;

    int16 public Karma;
    mapping(address => int8) private voters;

    string public PostString;
    
    address private Owner;
    uint256 public Balance;

    uint public CountComments;
    Comment[] public Comments;

    modifier isOwner(){
        require(tx.origin == Owner);
        _;
    }

    struct Comment{
        uint timestamp;
        uint CommentNumber;

        int16 Karma;
        mapping(address => int8) voters;

        string CommentString;

        address Poster;
    }

    constructor(string memory postString, uint CountPosts){
        timestamp = block.timestamp;
        PostString = postString;
        PostNumber = CountPosts;
        Owner = tx.origin;
    } 


    receive() payable external{
        Balance+=msg.value;
    }

    function AddComment(string memory CommentString) public {
        require(bytes(CommentString).length < 250, "Post exceeds character limit");
        Comment storage temp = Comments[CountComments++];

        temp.timestamp = block.timestamp;
        temp.CommentNumber = CountComments;

        temp.CommentString = CommentString;
        temp.Poster = msg.sender;
    }

    // Can't wait to get generics/templates or interfaces
    function Vote(int8 vote) public {
        require(voters[msg.sender] == 0 || vote == 0);
        if (vote != 0){
            Karma+=vote;
            voters[msg.sender]=vote;
        } else{
            Karma-=voters[msg.sender];
            voters[msg.sender]=vote;
        }
    }

    function Vote(int8 vote, uint CommentNumber) public {
        require(Comments[CommentNumber].voters[msg.sender] == 0|| vote == 0);
        if (vote != 0){
            Comments[CommentNumber].Karma+=vote;
            Comments[CommentNumber].voters[msg.sender]=vote;
        } else{
            Comments[CommentNumber].Karma-=Comments[CommentNumber].voters[msg.sender];
            Comments[CommentNumber].voters[msg.sender]=vote;
        }
    }

    function Transfer() public isOwner{
        payable(Owner).transfer(Balance);
        Balance = 0;
    }

    function DeletePost() public isOwner{
        selfdestruct(payable(Owner));
    }

}