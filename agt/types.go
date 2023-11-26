package agt

import (
	"github.com/BaptisteBuvron/go-voting/comsoc"
	"time"
)

// Request for [POST] /new_ballot
type RequestNewBallot struct {
	Rule         string               `json:"rule"`
	Deadline     time.Time            `json:"deadline"`
	Voters       []string             `json:"voter-ids"`
	Alternatives int                  `json:"#alts"`
	TieBreak     []comsoc.Alternative `json:"tie-break"`
}

// Response for [POST] /new_ballot
type ResponseNewBallot struct {
	BallotID string `json:"ballot-id"`
}

// Request for [POST] /vote
type RequestVote struct {
	AgentID  string               `json:"agent-id"`
	BallotID string               `json:"ballot-id"`
	Prefs    []comsoc.Alternative `json:"prefs"`
	Options  []int                `json:"options"` //Optional
}

// Request for [POST] /result
type RequestResult struct {
	BallotID string `json:"ballot-id"`
}

// Response for [POST] /result
type ResponseResult struct {
	Winner  comsoc.Alternative   `json:"winner"`
	Ranking []comsoc.Alternative `json:"ranking"`
}

// Generic response for message as error or /vote
type ResponseMessage struct {
	Message string `json:"message"`
}
