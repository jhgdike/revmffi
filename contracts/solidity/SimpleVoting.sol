// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleVoting {
    struct Proposal {
        string description;
        uint256 voteCount;
        bool executed;
    }

    struct Voter {
        bool hasVoted;
        uint256 votedProposal;
    }

    address public chairperson;
    mapping(address => Voter) public voters;
    Proposal[] public proposals;

    event ProposalCreated(uint256 proposalId, string description);
    event VoteCast(address voter, uint256 proposalId);
    event ProposalExecuted(uint256 proposalId);

    modifier onlyChairperson() {
        require(msg.sender == chairperson, "Only chairperson can call this function");
        _;
    }

    modifier hasNotVoted() {
        require(!voters[msg.sender].hasVoted, "Already voted");
        _;
    }

    constructor() {
        chairperson = msg.sender;
    }

    function createProposal(string memory _description) public onlyChairperson {
        proposals.push(Proposal({
            description: _description,
            voteCount: 0,
            executed: false
        }));
        emit ProposalCreated(proposals.length - 1, _description);
    }

    function vote(uint256 _proposalId) public hasNotVoted {
        require(_proposalId < proposals.length, "Invalid proposal");
        require(!proposals[_proposalId].executed, "Proposal already executed");

        Voter storage sender = voters[msg.sender];
        sender.hasVoted = true;
        sender.votedProposal = _proposalId;
        proposals[_proposalId].voteCount += 1;

        emit VoteCast(msg.sender, _proposalId);
    }

    function executeProposal(uint256 _proposalId) public onlyChairperson {
        require(_proposalId < proposals.length, "Invalid proposal");
        require(!proposals[_proposalId].executed, "Proposal already executed");

        proposals[_proposalId].executed = true;
        emit ProposalExecuted(_proposalId);
    }

    function getProposalCount() public view returns (uint256) {
        return proposals.length;
    }

    function getProposal(uint256 _proposalId) public view returns (
        string memory description,
        uint256 voteCount,
        bool executed
    ) {
        require(_proposalId < proposals.length, "Invalid proposal");
        Proposal storage proposal = proposals[_proposalId];
        return (proposal.description, proposal.voteCount, proposal.executed);
    }
} 