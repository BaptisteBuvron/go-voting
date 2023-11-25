package comsoc

import (
	"testing"
)

func TestCopelandSWF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}
	expectedCount := Count{
		0: 2,
		1: 1,
		2: 0,
	}
	count, err := CopelandSWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)
}
