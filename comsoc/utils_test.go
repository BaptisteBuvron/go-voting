package comsoc

import (
	"testing"
)

func TestCheckProfile(t *testing.T) {
	assert := Assert{t}

	// Should succeed
	prefs1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{3, 1, 2},
	}
	err := CheckProfile(prefs1)
	assert.NoError(err)

	// Should fail
	prefs2 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{3, 3, 4},
	}
	err = CheckProfile(prefs2)
	assert.Error(err)

	//Should fail
	prefs3 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{3, 1, 4},
	}
	err = CheckProfile(prefs3)
	assert.Error(err)

	// Should fail
	prefs4 := Profile{
		[]Alternative{1, 2, 4},
	}
	err = CheckProfile(prefs4)
	assert.Error(err)
}
