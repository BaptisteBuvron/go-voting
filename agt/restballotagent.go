package agt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type RestServerAgent struct {
	sync.Mutex
	id       string
	reqCount int
	addr     string
	ballots  map[string]*BallotAgent
}

func NewRestServerAgent(addr string) *RestServerAgent {
	return &RestServerAgent{id: addr, addr: addr}
}

// Test de la méthode
func checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

type Request struct {
	Operator string `json:"op"`
	Args     [2]int `json:"args"`
}

func respondJSON(w http.ResponseWriter, statuscode int, value any) {
	w.WriteHeader(statuscode)
	w.Header().Set("Content-Type", "application/json")
	serial, _ := json.Marshal(value)
	w.Write(serial)
}

type Response = func(statuscode int, value any)

func route[Request any](method string, do func(Request, Response)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if method is valid
		if !checkMethod(method, w, r) {
			respondJSON(w, http.StatusBadRequest, ResponseMessage{Message: fmt.Sprintf("Route support only %s", method)})
			return
		}
		// Deserialize json
		var request Request
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		err := json.Unmarshal(buf.Bytes(), &request)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, ResponseMessage{Message: "Invalid body"})
			return
		}
		// Run function
		do(request, func(statuscode int, value any) {
			respondJSON(w, statuscode, value)
		})
	}
}

func (rsa *RestServerAgent) doNewBallot(req RequestNewBallot, res Response) {
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++
	voteName := fmt.Sprintf("vote%d", rsa.reqCount)
	ballotAgent := NewBallotAgent(req)
	rsa.ballots[voteName] = ballotAgent
	responseNewBallot := ResponseNewBallot{BallotID: voteName}
	res(http.StatusCreated, responseNewBallot)
}

func (rsa *RestServerAgent) doVote(req RequestVote, res Response) {
	res(http.StatusOK, RequestVote{})
}

func (rsa *RestServerAgent) doResult(req RequestResult, res Response) {
	res(http.StatusOK, ResponseMessage{Message: ""})
}

func (rsa *RestServerAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", route("POST", rsa.doNewBallot))
	mux.HandleFunc("/vote", route("POST", rsa.doVote))
	mux.HandleFunc("/result", route("POST", rsa.doResult))

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
