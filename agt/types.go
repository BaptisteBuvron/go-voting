package agt

type RequestNewBallot struct {
	Rule         string   `json:"rule"`
	Deadline     string   `json:"deadline"`
	Voters       []string `json:"voter-ids"`
	Alternatives int      `json:"#alts"`
}

type ResponseNewBallot struct {
	BallotID string `json:"ballot-id"`
}

type RequestVote struct {
	AgentID string `json:"agent-id"`
	VoteID  string `json:"vote-id"`
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
