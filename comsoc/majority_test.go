package comsoc

import (
	"testing"
)

func TestMajoritySWF(t *testing.T) {
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
	count, err := MajoritySWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{1}
	bestAlts, err := MajoritySCF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, expectedBestAlts)

	// Test case  Empty profile
	_, emptyErr := MajoritySWF(Profile{})
	assert.NoError(emptyErr)

	// Test case equality
	equalityProfile := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{3, 1, 2},
	}
	expectedEqualityCount := Count{
		1: 1,
		2: 1,
		3: 1,
	}

	equalityCount, equalityErr := MajoritySWF(equalityProfile)
	assert.NoError(equalityErr)
	assert.DeepEqual(equalityCount, expectedEqualityCount)
}

func TestMajoritySCF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 3, 1},
		[]Alternative{1, 3, 2},
	}

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{1}
	bestAlts, err := MajoritySCF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, expectedBestAlts)
}
