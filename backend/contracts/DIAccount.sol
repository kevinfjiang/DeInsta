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
    
    address[] public Followers;
    mapping(address=>uint) private followersIndex;
    

    address[] public Following;
    mapping(address=>uint) private followingIndex;
    

    address private owner;
    
    constructor(string memory Name_) {
        Name = Name_;
        CountPosts = 0;
        owner = msg.sender;
        Followers.push(msg.sender);
        Following.push(msg.sender);
    }
    
    modifier isAdmin(){
        require(msg.sender == owner);
        _;
    }

    modifier isDIAccount(){
        require(DIAccount(msg.sender).Followers(0)==msg.sender);
        _;
    }
    

    function SetBio(string memory Bio_) external isAdmin {
       Bio=Bio_;
    }

    function DeletePost(uint PostID) external isAdmin{
        delete Posts[PostID];
    }
    
    function DeleteAccount() external payable isAdmin {
        selfdestruct(payable(owner));
    }    

    //Post features
    struct Post{
        uint timestamp;
        address Account;
        uint PostNumber;

        int16 Karma;
        mapping(address => int8) voters;

        string IPFSurl;
        string Caption;

        uint CountComments;
        Comment[] Comments;
        mapping(address => int8)[] CommentVoters;
    }

    function GetComments(uint PostID) external view returns(Comment[] memory CommentList){
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
    function CreatePost(string memory PostURL_, string memory Caption_) external isAdmin{
        require(bytes(Caption_).length < 50, "Post exceeds character limit"); 
        Post storage newPost = Posts[CountPosts];
        
        newPost.timestamp = block.timestamp;
        newPost.Account = address(this);
        newPost.PostNumber = CountPosts;

        newPost.IPFSurl = PostURL_;
        newPost.Caption = Caption_;

        emit PostBanner(newPost.timestamp, address(this), CountPosts);
        CountPosts++;
    }

    function Vote(address targ, int8 vote, uint PostNumber) external {
        vote = vote>0? int8(1) : (vote<0? -1:int8(0));
        require(Posts[PostNumber].voters[msg.sender] != vote, "You've already voted");
        DIAccount(targ).VoteOwner(vote, PostNumber);
    }

    // Can't wait to get generics/templates or interfaces
    function VoteOwner(int8 vote, uint PostNumber) external isDIAccount{
        Posts[PostNumber].Karma-=Posts[PostNumber].voters[msg.sender];
        Karma-=Posts[PostNumber].voters[msg.sender];
        Posts[PostNumber].Karma+=vote;
        Karma+=vote;
        Posts[PostNumber].voters[msg.sender]=vote;
    }


    function AddComment(address targ, uint PostNumber, string memory CommentString) external {
        require(bytes(CommentString).length < 50, "Post exceeds character limit");
        DIAccount(targ).AddCommentOwner(PostNumber, CommentString);
    }

    function AddCommentOwner(uint PostNumber, string memory CommentString) external isDIAccount{
        Comment storage newComment = Posts[PostNumber].Comments[Posts[PostNumber].CountComments];

        newComment.timestamp = block.timestamp;
        newComment.CommentNumber = Posts[PostNumber].CountComments;

        newComment.CommentString = CommentString;
        newComment.Poster = msg.sender;

        Posts[PostNumber].CountComments++;
    }

    function CommentVote(address targ, int8 vote, uint PostNumber, uint CommentNumber) external {
        vote = vote>0? int8(1): (vote<0? -1: int8(0));
        require(Posts[PostNumber].CommentVoters[CommentNumber][msg.sender] != vote, "You've already voted");
        DIAccount(targ).CommentVoteOwner(vote, PostNumber, CommentNumber);
    }

    function CommentVoteOwner(int8 vote, uint PostNumber, uint CommentNumber) external {
        Posts[PostNumber].Comments[CommentNumber].Karma-=Posts[PostNumber].CommentVoters[CommentNumber][msg.sender];
        Posts[PostNumber].Comments[CommentNumber].Karma+=vote;
        Posts[PostNumber].CommentVoters[CommentNumber][msg.sender]=vote;
    }

    function Follow(address targ) external {
        require(followersIndex[targ]==0 && DIAccount(targ).NewFollowing());
        Followers.push(targ);
        followersIndex[targ] = Followers.length-1;
    }

    function UnFollow(address targ) external {
        require(followersIndex[targ]!=0 && DIAccount(targ).UnNewFollowing());
        Followers[followersIndex[targ]] = Followers[Followers.length - 1];
        Followers.pop();
        followingIndex[targ] = 0;
    }

    function NewFollowing() external isDIAccount returns(bool added){
        require(followingIndex[msg.sender]==0);
        Following.push(msg.sender);
        followingIndex[msg.sender] = Followers.length-1;
        return true;
    }

    function UnNewFollowing() external isDIAccount returns(bool removed){
        require(followingIndex[msg.sender]!=0);
        Following.push(msg.sender);
        followingIndex[msg.sender] = Followers.length-1;
        return true;
    }

}

