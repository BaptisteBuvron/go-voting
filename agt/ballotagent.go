package agt

import (
	"fmt"
	"time"
	"tp3/agt"
	"tp3/comsoc"
)

// https://en.wikipedia.org/wiki/Ballot
type BallotAgent struct {
	rule 	 string
	deadline   time.Time
	voters	 []string
	alternativesNb int
	votted	 []string
	profiles   comsoc.Profile
}

type BallotAgentI interface {
	AgentI
}

func NewBallotAgent(request RequestNewBallot) *BallotAgent {
	return &BallotAgent{request.Rule, time.Now(), request.Voters, request.Alternatives, nil, nil}
}



func (b *BallotAgent) addVoter(vote RequestVote) (err error) {
	//verif if voter exists in voters
	exists := false
	for _, v := range b.voters {
		if v == vote.AgentID {
			exists = true
		}
	}
	if !exists {
		return agt.Error{agt.ErrorVoterNotFound, fmt.Sprintf("Voter %s not found", vote.AgentID)}
	}
	//verif if voter already voted
	for _, v := range b.votted {
		if v == vote.AgentID {
			return agt.Error{agt.ErrorAlreadyVoted, fmt.Sprintf("Voter %s already voted", vote.AgentID)}
		}
	}
	//verify the deadline
	if time.Now().After(b.deadline) {
		return agt.Error{agt.ErrorDeadline, fmt.Sprintf("Deadline %s is passed", b.deadline)}
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