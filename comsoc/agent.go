package comsoc

import (
	"fmt"
	"time"
)

type AgentID int

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative) bool
	Start()
}

type BallotAgent struct {
	Agent
	c          chan []Alternative
	voterCount int
	scf        func(Profile) (Alternative, error)
}

type BallotAgentI interface {
	AgentI
}

func (b *BallotAgent) Start() {
	go func() {
		profile := Profile{}
		for i := 0; i < b.voterCount; i++ {
			vote := <-b.c
			profile = append(profile, vote)
		}
		winner, err := b.scf(profile)
		if err != nil {
			fmt.Printf("Aucun vainqueur")
		} else {
			fmt.Printf("Vainqueur %d", winner)
		}
	}()
}

type VoterAgent struct {
	Agent
	c chan []Alternative
}

type VoterAgentI interface {
	AgentI
}

func (v *VoterAgent) Start() {
	go func() {
		v.c <- v.Prefs
	}()
}

func InitSystemDeVote(profile Profile, func(Profile) (Alternative, error)) {
	
}

// https://en.wikipedia.org/wiki/Ballot
func RunSystemDeVote(voters []VoterAgentI, ballot BallotAgentI) {
	ballot.Start()
	for _, voter := range voters {
		voter.Start()
	}
	time.Sleep(time.Minute)
}
