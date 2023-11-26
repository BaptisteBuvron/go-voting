package comsoc

import (
	"testing"
)

func TestCopelandSWF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{1, 3, 2},
	}
	expectedCount := Count{
		1: 2,
		2: 1,
		3: 0,
	}
	count, err := CopelandSWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case  Empty profile
	_, emptyErr := CopelandSWF(Profile{})
	assert.NoError(emptyErr)
}
