pragma solidity ^0.8.0;
pragma abicoder v2;
// SPDX-License-Identifier: MIT

contract DIAccount {
    string public Name;
    string public Bio;

    Post[] public Posts;
    int16 public Karma; // Add more robust Karma system
    
    // total posts
    uint public CountPosts;
    
    address owner;
    
    constructor(string memory Name_) {
        Name = Name_;
        CountPosts = 0;
        owner = msg.sender;
    }
    
    modifier isAdmin(){
        require(msg.sender == owner);
        _;
    }
    

    function SetBio(string memory Bio_) public isAdmin {
       Bio=Bio_;
    }

    function DeletePost(uint PostID) external isAdmin{
        delete Posts[PostID];
    }
    
    function DeleteAccount() public payable isAdmin {
        selfdestruct(payable(owner));
    }    

    //Post features
    struct Post{
        uint timestamp;
        uint PostNumber;

        int16 Karma;
        mapping(address => int8) voters;

        string IPFSurl;
        string Caption;

        uint CountComments;
        Comment[] Comments;
        mapping(address => int8)[] CommentVoters;
    }

    function GetComments(uint PostID) public view returns(Comment[] memory CommentList){
        return Posts[PostID].Comments;
    }

    struct Comment{
        uint timestamp;
        uint CommentNumber;

        int16 Karma;

        string CommentString;

        address Poster;
    }

    event PostBanner(
        uint indexed timestamp,
        address indexed owner,
        uint indexed PostNumber);

    // create new post
    function CreatePost(string memory PostURL_, string memory Caption_) public isAdmin{
        require(bytes(Caption_).length < 50, "Post exceeds character limit"); 
        Post storage newPost = Posts[CountPosts];
        
        newPost.timestamp = block.timestamp;
        newPost.PostNumber = CountPosts;

        newPost.IPFSurl = PostURL_;
        newPost.Caption = Caption_;

        emit PostBanner(newPost.timestamp, address(this), CountPosts);
        CountPosts++;
    }

    function GetPost(uint PostID) public returns(Comment[] memory){
        
    }

    // Can't wait to get generics/templates or interfaces
    function Vote(int8 vote, uint PostNumber) public {
        vote = vote>0? int8(1) : (vote<0? -1:int8(0));
        require(Posts[PostNumber].voters[msg.sender] != vote, "You've already voted");
        
        Posts[PostNumber].Karma-=Posts[PostNumber].voters[msg.sender];
        Karma-=Posts[PostNumber].voters[msg.sender];
        Posts[PostNumber].Karma+=vote;
        Karma+=vote;
        Posts[PostNumber].voters[msg.sender]=vote;
    }


    function AddComment(uint PostNumber, string memory CommentString) public {
        require(bytes(CommentString).length < 50, "Post exceeds character limit");
        Comment storage newComment = Posts[PostNumber].Comments[Posts[PostNumber].CountComments];

        newComment.timestamp = block.timestamp;
        newComment.CommentNumber = Posts[PostNumber].CountComments;

        newComment.CommentString = CommentString;
        newComment.Poster = msg.sender;

        Posts[PostNumber].CountComments++;
    }

    function CommentVote(int8 vote, uint PostNumber, uint CommentNumber) public {
        vote = vote>0? int8(1): (vote<0? -1: int8(0));
        require(Posts[PostNumber].CommentVoters[CommentNumber][msg.sender] != vote, "You've already voted");
        
        Posts[PostNumber].Comments[CommentNumber].Karma-=Posts[PostNumber].CommentVoters[CommentNumber][msg.sender];
        Posts[PostNumber].Comments[CommentNumber].Karma+=vote;
        Posts[PostNumber].CommentVoters[CommentNumber][msg.sender]=vote;
    }


}