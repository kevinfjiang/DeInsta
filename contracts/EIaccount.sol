pragma solidity ^0.8.0;
// SPDX-License-Identifier: MIT

contract EIaccount { // Following
    string public Name;
    string public Bio;

    Post[] public Posts;
    int16 public Karma; // Add more robust Karma system
    
    // total posts
    uint public CountPosts;
    
    address owner;
    mapping(address => bool) public Followers;
    uint public CountFollowers;
    address[] public Following;
    uint public CountFollowing;
    
    struct Post{
        uint timestamp;
        uint PostNumber;

        int16  Karma;
        mapping(address => int8) voters;

        string Hash;
        string  PostDescription;
        
        address Owner;

        uint CountComments;
        mapping(uint => Comment) Comments;
    }

    event PostBanner(
        address User,
        uint id,

        string  Hash,
        string  PostDescription);

    // create new post
    function AddPost(string memory Hash_, string memory Description_) public isAdmin {
        require(bytes(Hash_).length > 0);
        require(bytes(Description_).length < 100, "Post exceeds character limit");
        require(msg.sender!=address(0));
        
        Post storage newPost = Posts[CountPosts];

        newPost.timestamp = block.timestamp;
        newPost.Hash= Hash_;
        newPost.PostDescription = Description_;
        newPost.PostNumber = CountPosts;
        newPost.Owner = msg.sender;
        
        emit PostBanner(address(this), CountPosts, Hash_, Description_);
        CountPosts++;
        
    }

    struct Comment{
        uint timestamp;
        uint CommentNumber;

        int16 Karma;
        mapping(address => int8) voters;

        string CommentString;

        address Poster;
    }

    function AddComment(string memory CommentString, uint PostNum) public {
        require(bytes(CommentString).length < 100, "Post exceeds character limit");
        Comment storage newComment = Posts[PostNum].Comments[Posts[PostNum].CountComments];

        newComment.timestamp = block.timestamp;
        newComment.CommentNumber = Posts[PostNum].CountComments;

        newComment.CommentString = CommentString;
        newComment.Poster = msg.sender;
        
        Posts[PostNum].CountComments++;
    }

    // Can't wait to get generics/templates or interfaces
    function Vote(int8 vote, uint PostNum) public {
        vote = vote>0? int8(1) : (vote<0? -1:int8(0));
        require(Posts[PostNum].voters[msg.sender] != vote, "Must change prev vote");
        
        Karma-=Posts[PostNum].voters[msg.sender];
        Posts[PostNum].Karma-=Posts[PostNum].voters[msg.sender];
        Karma+=vote;
        Posts[PostNum].Karma+=vote;
        Posts[PostNum].voters[msg.sender]=vote;
    }


    constructor(string memory Name_, string memory Bio_) {
        require(bytes(Name_).length < 15, "Name exceeds character limit");
        require(bytes(Bio_).length < 30, "Bio exceeds character limit");
        Name = Name_;
        Bio = Bio_;
        CountPosts = 0;
        owner = msg.sender;
    }
    
    modifier isAdmin(){
        require(msg.sender == owner);
        _;
    }
    

    function DeletePost(uint PostID) external isAdmin{
        delete Posts[PostID];
    }
    
    function GetOwnerAddress() public view returns (address adminAddress) {
        return owner;
    }
    
    function Transfer() public isAdmin{
        payable(owner).transfer(address(this).balance);
    }

    function adminDeleteAccount() public payable isAdmin {
        selfdestruct(payable(owner));
    }    

    
}