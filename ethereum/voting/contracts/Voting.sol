// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;

contract Voting {
    struct voter { //选民
        address voterAddress;
        uint tokenNum;
        uint[] tokensVoteForCandidates;
    }

    uint public totalTokens;
    uint public tokenBalance;
    uint public tokenPrice;

    bytes32[] public candidateList; //候选人列表

    mapping(bytes32=>uint) public votesReceived;
    mapping(address=>voter) public voterInfo;

    constructor(uint totalSupply, uint price, bytes32[] memory candidateNames) {
        totalTokens = totalSupply;
        tokenBalance = totalSupply;
        tokenPrice = price;
        candidateList = candidateNames;
    }

    function buy() payable public returns(uint) {
        uint tokensToBuy = msg.value / tokenPrice;
        require(tokensToBuy <= tokenBalance);
        voterInfo[msg.sender].voterAddress = msg.sender;
        voterInfo[msg.sender].tokenNum += tokensToBuy;
        tokenBalance -= tokensToBuy;
        return tokensToBuy;
    }

    function voteForCandidate(bytes32 candidate, uint voteTokens) public {
        int index = indexOfCandidate(candidate);
        require(index != -1);
        if (voterInfo[msg.sender].tokensVoteForCandidates.length == 0) {
            for(uint i = 0; i < candidateList.length; i++) {
                voterInfo[msg.sender].tokensVoteForCandidates.push(0);
            }
        }
        uint availableTokens = voterInfo[msg.sender].tokenNum - totalUsedTokens(voterInfo[msg.sender].tokensVoteForCandidates);
        require(availableTokens >= voteTokens);
        votesReceived[candidate] += voteTokens;
        voterInfo[msg.sender].tokenNum -= voteTokens;
        voterInfo[msg.sender].tokensVoteForCandidates[uint(index)] += voteTokens;
    }

    function totalVotesFor(bytes32 candidate) public view returns(uint) {
        return votesReceived[candidate];
    }

    function totalUsedTokens(uint[] memory votesForCandidate) public pure returns(uint) {
        uint useTokens = 0;
        for(uint i = 0; i < votesForCandidate.length; i++){
            useTokens += votesForCandidate[i];
        }
        return useTokens;
    }

    function indexOfCandidate(bytes32 candidate) public view returns(int) {
        for(uint i = 0; i < candidateList.length; i++) {
            if(candidate == candidateList[i])
                return int(i);
        }
        return -1;
    }

    function tokenSold() public view returns(uint) {
        return totalTokens - tokenBalance;
    }

    function voterDetails(address voterAddr) public view returns(uint, uint[] memory) {
        return (voterInfo[voterAddr].tokenNum, voterInfo[voterAddr].tokensVoteForCandidates);
    }

    function allCandidate() public view returns(bytes32[] memory) {
        return candidateList;
    }

    function transfer(address payable _to) public {
        _to.transfer(address(this).balance);
    }
}