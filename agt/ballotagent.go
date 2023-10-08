package agt

import (
	"fmt"
	"tp3/comsoc"
)

// https://en.wikipedia.org/wiki/Ballot
type BallotAgent struct {
	Agent
	Channel    chan []comsoc.Alternative
	VoterCount int
	Scf        func(comsoc.Profile) (comsoc.Alternative, error)
}

type BallotAgentI interface {
	AgentI
}

func (b *BallotAgent) Start() {
	go func() {
		profile := comsoc.Profile{}
		for i := 0; i < b.VoterCount; i++ {
			vote := <-b.Channel
			profile = append(profile, vote)
		}
		fmt.Printf("[%-10s] Profile %d\n", b.Name, profile)
		winner, err := b.Scf(profile)
		if err != nil {
			fmt.Printf("[%-10s] Aucun vainqueur\n", b.Name)
		} else {
			fmt.Printf("[%-10s] Vainqueur %d\n", b.Name, winner)
		}
	}()
}
