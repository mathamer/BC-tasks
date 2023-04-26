// Create a Solidity contract that implements a simple voting system. The contract should have the following functionality:
// Allows users to vote for a candidate by calling vote() function.
// Keep track of the number of votes for each candidate.
//Prevent users from voting twice.

// make modifiers to enforce the following rules:
//the vote(uint broj) function should be restricted to users who have not voted before.
// the voitng can be closed by the owner of the contract by calling the closeVoting() function.

pragma solidity >=0.8.2 <0.9.0;

contract Voting {
    uint broj;
    address private owner;
    bool private votingClosed;

    mapping (uint => uint) private votes;

    mapping (address => bool) private voted;

    constructor() {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
    }

    modifier hasVoted() {
        require(!voted[msg.sender], "Already voted");
        _;
    }

    function vote() public hasVoted votingNotClosed{
        voted[msg.sender] = true;
        votes[broj]++;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can close voting");
        _;
    }

    modifier votingNotClosed() {
        require(!votingClosed, "Voting is closed");
        _;
    }

    function closeVoting() public onlyOwner votingNotClosed{
        votingClosed = true;
    }

    function getVotes(uint broj) public view returns (uint) {
        return votes[broj];
    }

    function getVotingClosed() public view returns (bool) {
        return votingClosed;
    }
}