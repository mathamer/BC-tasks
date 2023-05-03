// https://hackmd.io/@rosaj/Bk-wqdy4n
// 2. Smart Contract for a Crowdfunding Platform
// Overview:
    // In this assignment, you will be creating a Solidity smart contract for a crowdfunding platform. 
    // The platform will allow users to create campaigns and contribute to them. 
    // The contract will keep track of the amount of funds raised for each campaign and allow the creator to withdraw funds once the campaign has ended successfully.

// Part 1: Campaign Struct
// Create a Campaign struct that contains the following fields:

// address payable creator: The Ethereum address of the campaign creator.
// string title: The title of the campaign.
// uint256 goal: The funding goal of the campaign.
// uint256 raisedAmount: The amount of funds raised so far.
// uint256 deadline: The deadline for the campaign to end.

// Part 2: Campaign Contract

// Create a Campaign contract that contains the following functions and modifiers:

// Functions:
// createCampaign(string _title, uint256 _goal, uint256 _deadline): This function should create a new campaign with the given parameters and add it to an array of campaigns.
// contribute(uint256 _campaignIndex): This function should allow users to contribute to a campaign by sending ether to the contract. The amount contributed should be added to the raisedAmount field of the Campaign struct for the given campaign index.
// withdrawFunds(uint256 _campaignIndex): This function should allow the creator of a campaign to withdraw the raised funds if the campaign has met its funding goal and the deadline has passed.

// Modifiers:
// onlyCreator(uint256 _campaignIndex): This modifier should ensure that only the creator of a campaign can call certain functions, such as withdrawFunds().
// campaignExists(uint256 _campaignIndex): This modifier should ensure that the campaign index being passed to a function corresponds to an existing campaign.

pragma solidity >=0.8.2 <0.9.0;

contract Crowfunding {
    struct Campaign {
        address payable creator;
        string title;
        uint256 goal;
        uint256 raisedAmount;
        uint256 deadline;
    }

    Campaign[] private campaigns;

    modifier onlyCreator(uint256 _campaignIndex) {
        require(msg.sender == campaigns[_campaignIndex].creator, "Only creator can withdraw funds");
        _;
    }

    modifier campaignExists(uint256 _campaignIndex) {
        require(_campaignIndex < campaigns.length, "Campaign does not exist");
        _;
    }

    function createCampaign(string memory _title, uint256 _goal, uint256 _deadline) public {
        campaigns.push(Campaign(payable(msg.sender), _title, _goal, 0, _deadline));
    }

    function contribute(uint256 _campaignIndex) public payable campaignExists(_campaignIndex) {
        campaigns[_campaignIndex].raisedAmount += msg.value;
    }

    function withdrawFunds(uint256 _campaignIndex) public onlyCreator(_campaignIndex) campaignExists(_campaignIndex) hasMetFundingGoal(_campaignIndex) hasEnded(_campaignIndex){
        payable(campaigns[_campaignIndex].creator).transfer(campaigns[_campaignIndex].raisedAmount);
    }

    modifier hasMetFundingGoal(uint256 _campaignIndex) {
        require(campaigns[_campaignIndex].raisedAmount >= campaigns[_campaignIndex].goal, "Campaign has not met its funding goal");
        _;
    }

    modifier hasEnded(uint256 _campaignIndex) {
        require(block.timestamp >= campaigns[_campaignIndex].deadline, "Campaign has not ended yet");
        _;
    }
}