package agt

import (
	"fmt"
	"time"
)

type RequestNewBallot struct {
	Rule         string   `json:"rule"`
	Deadline     string   `json:"deadline"`
	Voters       []string `json:"voter-ids"`
	Alternatives int      `json:"#alts"`
	TieBreak     []int    `json:"tie-break"`
}

func isRequestNewBallotValid(req RequestNewBallot) (bool, error) {
	// Check if the Rule is set
	if req.Rule == "" {
		return false, fmt.Errorf("Rule is not set")
	}

	// Check if the Deadline is set
	if req.Deadline == "" {
		return false, fmt.Errorf("Deadline is not set")
	}

	//check if the deadline is valid (RFC3339)
	_, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		return false, fmt.Errorf("Invalid deadline")
	}

	// Check if at least one voter is present
	if len(req.Voters) == 0 {
		return false, fmt.Errorf("No voters specified")
	}

	// Check if the number of alternatives is valid
	if req.Alternatives <= 0 {
		return false, fmt.Errorf("Invalid number of alternatives")
	}

	// Check if TieBreak is set and has at least one element
	if len(req.TieBreak) == 0 {
		return false, fmt.Errorf("TieBreak is not set")
	}

	return true, nil
}

type ResponseNewBallot struct {
	BallotID string `json:"ballot-id"`
}

type RequestVote struct {
	AgentID string `json:"agent-id"`
	VoteID  string `json:"ballot-id"`
	Prefs   []int  `json:"prefs"`
	Options []int  `json:"options"` //Optional
}

func isRequestVoteValid(req RequestVote) (bool, error) {
	// Check if the AgentID is set
	if req.AgentID == "" {
		return false, fmt.Errorf("AgentID is not set")
	}

	// Check if the VoteID is set
	if req.VoteID == "" {
		return false, fmt.Errorf("VoteID is not set")
	}

	// Check if the Prefs is set
	if len(req.Prefs) == 0 {
		return false, fmt.Errorf("Prefs is not set")
	}

	// Options is optional

	return true, nil
}

type RequestResult struct {
	BallotID string `json:"ballot-id"`
}

func isRequestResultValid(req RequestResult) (bool, error) {
	// Check if the BallotID is set
	if req.BallotID == "" {
		return false, fmt.Errorf("BallotID is not set")
	}

	return true, nil
}

type ResponseResult struct {
	Winner  string `json:"winner"`
	Ranking []int  `json:"ranking"`
}

type ResponseMessage struct {
	Message string `json:"error"`
}
