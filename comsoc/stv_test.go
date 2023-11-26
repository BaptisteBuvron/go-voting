package comsoc

import (
	"testing"
)

func TestSTV_SWF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: 1 have majority at 1st round
	profile := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{1, 3, 2},
	}
	expectedCount := Count{
		1: 1,
		2: 0,
		3: 0,
	}
	count, err := STV_SWF(profile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case 2: equality at first turn then 2 win
	profile = Profile{
		[]Alternative{3, 2, 1},
		[]Alternative{2, 3, 1},
		[]Alternative{1, 3, 2},
	}
	expectedCount = Count{
		1: 1,
		2: 2,
		3: 0,
	}
	count, err = STV_SWF(profile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)
}
