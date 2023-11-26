package agt

import (
	"ia04/comsoc"
	"time"
)

// Agent used for voting
type RestClientAgent struct {
	url string
	id  string
}

// Create an new client agent
func NewRestClientAgent(url string, id string) *RestClientAgent {
	return &RestClientAgent{url, id}
}

// Create a ballot by sending a request to the RestBallotAgent
func (client *RestClientAgent) CreateBallot(rule string, deadline time.Time, voters []string, alternatives int, tieBreak []comsoc.Alternative) (string, error) {
	req := RequestNewBallot{rule, deadline, voters, alternatives, tieBreak}
	res, err := request[ResponseNewBallot](client.url+"/new_ballot", req)
	return res.BallotID, err
}

// Wait server to be available
func (client *RestClientAgent) WaitAvailable(duration time.Duration) bool {
	return WaitAvailable(client.url, duration)
}

// Vote by sending a request to the RestBallotAgent
func (client *RestClientAgent) Vote(ballotID string, prefs []comsoc.Alternative, options []int) error {
	_, err := request[ResponseMessage](client.url+"/vote", RequestVote{client.id, ballotID, prefs, options})
	return err
}

// Get the result of an ballot by sending a request to the RestBallotAgent
func (client *RestClientAgent) Result(ballotID string) (ResponseResult, error) {
	return request[ResponseResult](client.url+"/result", RequestResult{ballotID})
}
