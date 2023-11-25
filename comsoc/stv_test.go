package comsoc

import (
	"testing"
)

func TestSTV_SWF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: 0 have majority at 1st round
	profile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}
	expectedCount := Count{
		0: 1,
		1: 0,
		2: 0,
	}
	count, err := STV_SWF(profile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case 2: equality at first turn then 1 win
	profile = Profile{
		[]Alternative{2, 1, 0},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}
	expectedCount = Count{
		0: 1,
		1: 2,
		2: 0,
	}
	count, err = STV_SWF(profile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)
}
