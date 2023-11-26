package agt

import (
	"fmt"
	"ia04/comsoc"
	"os"
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
func (client *RestClientAgent) CreateBallot(rule string, deadline time.Time, voters []string, alternatives int, tieBreak []comsoc.Alternative) {
	req := RequestNewBallot{rule, deadline, voters, alternatives, tieBreak}
	res, err := request[ResponseNewBallot](client.url+"/new_ballot", req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	} else {
		fmt.Println(res.BallotID)
	}
}

// Vote by sending a request to the RestBallotAgent
func (client *RestClientAgent) Vote(ballotID string, prefs []comsoc.Alternative, options []int) {
	req := RequestVote{client.id, ballotID, prefs, options}
	_, err := request[ResponseMessage](client.url+"/vote", req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

// Get the result of an ballot by sending a request to the RestBallotAgent
func (client *RestClientAgent) Result(ballotID string) {
	res, err := request[ResponseResult](client.url+"/result", RequestResult{ballotID})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	} else {
		fmt.Printf("Winner: %v\nRanking: %v\n", res.Winner, res.Ranking)
	}
}
