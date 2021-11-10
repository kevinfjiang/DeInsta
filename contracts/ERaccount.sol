pragma solidity ^0.8.0;
// SPDX-License-Identifier: MIT

contract ERAccount {
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
    function PostMessage(string memory PostName_, string memory PostString_) public isAdmin returns(address postaddress) {
        require(bytes(PostString_).length < 500, "Post exceeds character limit"); 
        Post TempPost = new Post(PostName_, PostString_, CountPosts); 
        Posts.push(address(TempPost));
        return address(TempPost);
    }

    function DeletePost(uint PostID) external isAdmin{
        Post(Posts[PostID]).DeletePost();
        delete Posts[PostID];
    }

    function GetPost(uint PostID) public view returns(address postaddress){
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

    string public PostName;
    string public PostString;
    
    address private Owner;

    uint public CountComments;
    mapping(uint => Comment) public Comments;

    modifier isOwner(){
        require(msg.sender == Owner);
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

    event PostBanner(
        uint timestamp,
        int16 Karma,

        address Post,

        string PostName,
        string  PostString);

    constructor(string memory PostName_, string memory PostString_, uint CountPosts){
        timestamp = block.timestamp;
        PostName = PostName_;
        PostString = PostString_;
        PostNumber = CountPosts;
        Owner = tx.origin;
        emit PostBanner(timestamp, Karma, address(this), PostName, PostString);
    } 

    function AddComment(string memory CommentString) public {
        require(bytes(CommentString).length < 250, "Post exceeds character limit");
        Comment storage newComment = Comments[CountComments];

        newComment.timestamp = block.timestamp;
        newComment.CommentNumber = CountComments;

        newComment.CommentString = CommentString;
        newComment.Poster = msg.sender;
        
        CountComments++;
    }

    // Can't wait to get generics/templates or interfaces
    function Vote(int8 vote) public {
        vote = vote>0? int8(1) : (vote<0? -1:int8(0));
        require(voters[msg.sender] != vote, "Must change prev vote");
        
        Karma-=voters[msg.sender];
        Karma+=vote;
        voters[msg.sender]=vote;
        emit PostBanner(timestamp, Karma, address(this), PostName, PostString);
    }

    function Vote(int8 vote, uint CommentNumber) public {
        vote = vote>0? int8(1): (vote<0? -1: int8(0));
        require(Comments[CommentNumber].voters[msg.sender] != vote, "Must change prev vote");
        
        Comments[CommentNumber].Karma-=Comments[CommentNumber].voters[msg.sender];
        Comments[CommentNumber].Karma+=vote;
        Comments[CommentNumber].voters[msg.sender]=vote;
    }

    function Transfer() public isOwner{
        payable(Owner).transfer(address(this).balance);
    }

    function DeletePost() public isOwner{
        selfdestruct(payable(Owner));
    }

}