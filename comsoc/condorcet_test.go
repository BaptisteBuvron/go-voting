package comsoc

import (
	"testing"
)

func TestCondorcetWinner(t *testing.T) {
	assert := Assert{t}

	// 0 est le gagnant de Condorcet
	prefs1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{0, 2, 1},
		[]Alternative{2, 1, 0},
	}
	// 0 win against 2 and 1
	// 2 win against 1
	bestAlts, err := CondorcetWinner(prefs1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, Alts(0))

	// No Condorcet Winner
	prefs2 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 1},
	}
	bestAlts, err = CondorcetWinner(prefs2)
	assert.NoError(err)
	assert.Empty(bestAlts)

	// No Condorcet Winner
	prefs3 := Profile{
		[]Alternative{0, 2, 1},
		[]Alternative{0, 2, 1},
		[]Alternative{0, 2, 1},
		[]Alternative{2, 1, 0},
		[]Alternative{2, 1, 0},
		[]Alternative{1, 2, 0},
	}
	// 2 win against 1
	bestAlts, err = CondorcetWinner(prefs3)
	assert.NoError(err)
	assert.Empty(bestAlts)
}
