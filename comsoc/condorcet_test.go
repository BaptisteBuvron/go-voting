package comsoc

import (
	"testing"
)

func TestCondorcetWinner(t *testing.T) {
	assert := Assert{t}

	// 1 is Condorcet winner
	prefs1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{1, 3, 2},
		[]Alternative{3, 2, 1},
	}
	// 1 win against 3 and 2
	// 3 win against 2
	bestAlts, err := CondorcetWinner(prefs1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, []Alternative{1})

	// No Condorcet Winner
	prefs2 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{3, 1, 2},
	}
	bestAlts, err = CondorcetWinner(prefs2)
	assert.NoError(err)
	assert.Empty(bestAlts)

	// No Condorcet Winner
	prefs3 := Profile{
		[]Alternative{1, 3, 2},
		[]Alternative{1, 3, 2},
		[]Alternative{1, 3, 2},
		[]Alternative{3, 2, 1},
		[]Alternative{3, 2, 1},
		[]Alternative{2, 3, 1},
	}
	// 3 win against 2
	bestAlts, err = CondorcetWinner(prefs3)
	assert.NoError(err)
	assert.Empty(bestAlts)
}
