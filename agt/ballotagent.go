package agt

import (
	"fmt"
	"net/http"
	"time"
	"tp3/comsoc"
)

// https://en.wikipedia.org/wiki/Ballot
type BallotAgent struct {
	rule           string
	deadline       time.Time
	voters         []string
	alternativesNb int
	tieBreak       []int
	votted         []string
	profiles       comsoc.Profile
}

type BallotAgentI interface {
	AgentI
}

func NewBallotAgent(request RequestNewBallot) *BallotAgent {
	deadline, _ := time.Parse(time.RFC3339, request.Deadline)
	return &BallotAgent{request.Rule, deadline, request.Voters, request.Alternatives, request.TieBreak, nil, nil}
}

func (b *BallotAgent) addVoter(vote RequestVote) (err Error) {
	//verif if voter exists in voters
	exists := false
	for _, v := range b.voters {
		if v == vote.AgentID {
			exists = true
		}
	}
	if !exists {
		return Error{ErrorVoterNotFound, fmt.Sprintf("Voter %s not found", vote.AgentID)}
	}
	//verif if voter already voted
	for _, v := range b.votted {
		if v == vote.AgentID {
			return Error{ErrorAlreadyVoted, fmt.Sprintf("Voter %s already voted", vote.AgentID)}
		}
	}
	//verify the deadline
	if time.Now().After(b.deadline) {
		return Error{ErrorDeadline, fmt.Sprintf("Deadline %s is passed", b.deadline)}
	}
	//verify if the number of alternatives is correct
	if len(vote.Prefs) != b.alternativesNb {
		return Error{http.StatusBadRequest, fmt.Sprintf("The number of alternatives is incorrect")}
	}
	b.votted = append(b.votted, vote.AgentID)
	// convert []int to []comsoc.Alternative
	prefs := make([]comsoc.Alternative, len(vote.Prefs))
	for i, pref := range vote.Prefs {
		prefs[i] = comsoc.Alternative(pref)
	}
	b.profiles = append(b.profiles, prefs)
	return
}
