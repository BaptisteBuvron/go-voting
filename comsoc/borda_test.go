package comsoc

import (
	"testing"
)

func TestBordaSWF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 1, 3},
		[]Alternative{3, 2, 1},
	}

	expectedCount := Count{
		1: 3,
		2: 4,
		3: 2,
	}

	count, err := BordaSWF(profile1)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case  Empty profile
	_, emptyErr := BordaSWF(Profile{})
	assert.NoError(emptyErr)
}

func TestBordaSCF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 1, 3},
		[]Alternative{3, 2, 1},
	}

	bestAlts, err := BordaSCF(profile1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, []Alternative{2})
}
