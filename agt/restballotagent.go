package agt

import (
	"ia04/comsoc"
	"log"
	"net/http"
	"sync"
	"time"
)

// Server Agent for create BallotAgent
type RestServerAgent struct {
	sync.Mutex
	id      string
	addr    string
	ballots map[string]*BallotAgent
}

// Create RestServerAgent
func NewRestServerAgent(addr string) *RestServerAgent {
	ballots := make(map[string]*BallotAgent)
	return &RestServerAgent{id: addr, addr: addr, ballots: ballots}
}

// [POST] /new_ballot
func (server *RestServerAgent) doNewBallot(req RequestNewBallot, res Response) error {
	// We need to lock for being thread-safe
	server.Lock()
	defer server.Unlock()

	// Try to create ballot
	ballotAgent, err := NewBallotAgent(req.Rule, req.Deadline, req.Voters, req.Alternatives, req.TieBreak)
	if err != nil {
		return err
	}

	// Save ballot to server
	server.ballots[ballotAgent.id] = ballotAgent

	// Reply for the ballot
	responseNewBallot := ResponseNewBallot{BallotID: ballotAgent.id}
	res(http.StatusCreated, responseNewBallot)
	return nil
}

// [POST] /vote
func (server *RestServerAgent) doVote(req RequestVote, res Response) error {
	// We need to lock for being thread-safe
	server.Lock()
	defer server.Unlock()

	// Get the ballot from server
	ballotAgent, ok := server.ballots[req.BallotID]
	if !ok {
		return comsoc.HTTPErrorf(http.StatusNotFound, "Ballot %s not found", req.BallotID)
	}

	// Add vote to ballot or raise error
	err := ballotAgent.Vote(req.AgentID, req.Prefs, req.Options)
	if err != nil {
		return err
	}

	// Send result
	res(http.StatusOK, ResponseMessage{Message: "OK"})
	return nil
}

// [POST] /result
func (server *RestServerAgent) doResult(req RequestResult, res Response) error {
	// We need to wait the server to be in unlocked phase (but we dont need to lock because we read-only)
	server.Lock()
	server.Unlock() // Unlock directly

	// Get the ballot from server
	ballotAgent, ok := server.ballots[req.BallotID]
	if !ok {
		return comsoc.HTTPErrorf(http.StatusNotFound, "Ballot %s not found", req.BallotID)
	}

	// Add vote to ballot or raise error
	winner, ranking, err := ballotAgent.result()
	if err != nil {
		return err
	}
	res(http.StatusOK, ResponseResult{Winner: winner, Ranking: ranking})
	return nil

}

func (server *RestServerAgent) Start() {
	// Creation of the multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", route("POST", server.doNewBallot))
	mux.HandleFunc("/vote", route("POST", server.doVote))
	mux.HandleFunc("/result", route("POST", server.doResult))

	// Creating the http server
	s := &http.Server{
		Addr:           server.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// Server launch
	log.Println("Listening on", server.addr)
	go log.Fatal(s.ListenAndServe())
}
