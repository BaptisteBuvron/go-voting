package agt

import (
	"fmt"
	"tp3/comsoc"
)

type VoterAgent struct {
	Agent
	Channel chan []comsoc.Alternative
}

type VoterAgentI interface {
	AgentI
}

func (v *VoterAgent) Start() {
	go func() {
		fmt.Printf("[%-10s] Je vote : %d\n", v.Name, v.Prefs)
		v.Channel <- v.Prefs
	}()
}
