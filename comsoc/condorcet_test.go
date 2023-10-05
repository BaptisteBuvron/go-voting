package comsoc

import (
	"fmt"
	"testing"
)

func TestCondorcetWinner(t *testing.T) {
	//0 est le gagnant de Condorcet
	prefs1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{0, 2, 1},
		[]Alternative{2, 1, 0},
	}
	bestAlts, err := CondorcetWinner(prefs1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(bestAlts)
	if len(bestAlts) != 1 {
		t.Errorf("Error: %s", err)
	}
	if bestAlts[0] != 0 {
		t.Errorf("Error: %s", err)
	}

	//Il n'y a pas de gagnant de Condorcet
	prefs2 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 1},
	}
	bestAlts, err = CondorcetWinner(prefs2)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(bestAlts)
	if len(bestAlts) != 0 {
		t.Errorf("Error: %s", err)
	}

	// Todo tester avec une égalité non strict
}
