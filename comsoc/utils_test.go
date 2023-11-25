package comsoc

import (
	"testing"
)

func TestCheckProfile(t *testing.T) {
	assert := Assert{t}

	// Should succeed
	prefs1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 1},
	}
	err := CheckProfile(prefs1)
	assert.NoError(err)

	// Should fail
	prefs2 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 2, 3},
	}
	err = CheckProfile(prefs2)
	assert.Error(err)

	//Should fail
	prefs3 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 3},
	}
	err = CheckProfile(prefs3)
	assert.Error(err)

	// Should fail (eeuuuhh non ?)
	prefs4 := Profile{
		[]Alternative{0, 1, 3},
	}
	err = CheckProfile(prefs4)
	assert.Error(err)
}
