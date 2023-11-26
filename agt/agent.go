package agt

import (
	"fmt"
	"ia04/comsoc"
	"reflect"
)

// ID for agent
type AgentID int

// Agent structure
type Agent struct {
	ID    AgentID
	Name  string
	Prefs []comsoc.Alternative
}

// Compare an agent based on his id
func (a *Agent) Equal(ag any) bool {
	otherAgent, ok := ag.(*Agent)
	if !ok {
		return false
	}
	return a.ID == otherAgent.ID
}

// Deeply compare and agent
func (a *Agent) DeepEqual(ag any) bool {
	otherAgent, ok := ag.(*Agent)
	if !ok {
		return false
	}
	return a.ID == otherAgent.ID && a.Name == otherAgent.Name && reflect.DeepEqual(a.Prefs, otherAgent.Prefers)
}

// Clone an agent
func (a *Agent) Clone() Agent {
	clone := *a
	return clone
}

// represent an agent
func (a *Agent) String() string {
	return fmt.Sprintf("<Agent ID: %d Name: %s Prefs: %d>", a.ID, a.Name, a.Prefs)
}

// Ask agent his preferences
func (ag *Agent) Prefers(a comsoc.Alternative, b comsoc.Alternative) bool {
	return comsoc.IsPref(a, b, ag.Prefs)
}

// Méthode Start
func (a *Agent) Start() {
	fmt.Printf("L'agent %s a démarré.\n", a.Name)
}
