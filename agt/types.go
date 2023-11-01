package agt

import "fmt"

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
	Options []int  `json:"options"`
}

type RequestResult struct {
	BallotID string `json:"ballot-id"`
}

type ResponseResult struct {
	Winner  string `json:"winner"`
	Ranking []int  `json:"ranking"`
}

type ResponseMessage struct {
	Message string `json:"error"`
}
