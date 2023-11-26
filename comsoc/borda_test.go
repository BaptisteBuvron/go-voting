package comsoc

import (
	"testing"
)

func TestBordaSWF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 0, 2},
		[]Alternative{2, 1, 0},
	}

	expectedCout := Count{
		0: 3,
		1: 4,
		2: 2,
	}

	count, err := BordaSWF(profile1)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCout)
}

func TestBordaSCF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 0, 2},
		[]Alternative{2, 1, 0},
	}

	bestAlts, err := BordaSCF(profile1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, []Alternative{1})
}
