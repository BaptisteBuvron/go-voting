package agt

import (
	"fmt"
	"ia04/comsoc"
)

type AgentID int

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []comsoc.Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a comsoc.Alternative, b comsoc.Alternative) bool
	Start()
}

// Méthode Equal
func (a *Agent) Equal(ag AgentI) bool {
	otherAgent, ok := ag.(*Agent)
	if !ok {
		return false
	}
	return a.ID == otherAgent.ID
}

func (a *Agent) DeepEqual(ag AgentI) bool {
	otherAgent, ok := ag.(*Agent)
	if !ok {
		return false
	}
	return a.ID == otherAgent.ID && a.Name == otherAgent.Name
}

func (a *Agent) Clone() AgentI {
	clone := *a
	return &clone
}

func (a *Agent) String() string {
	return fmt.Sprintf("<Agent ID: %d Name: %s Prefs: %d>", a.ID, a.Name, a.Prefs)
}

func (ag *Agent) Prefers(a comsoc.Alternative, b comsoc.Alternative) bool {
	return comsoc.IsPref(a, b, ag.Prefs)
}

// Méthode Start
func (a *Agent) Start() {
	fmt.Printf("L'agent %s a démarré.\n", a.Name)
}
