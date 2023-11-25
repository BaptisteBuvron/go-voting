package comsoc

import (
	"testing"
)

func TestMajoritySWF(t *testing.T) {
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
	count, err := MajoritySWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{0}
	bestAlts, err := MajoritySCF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, expectedBestAlts)

	// Test case  Empty profile
	emptyProfile := Profile{}
	_, emptyErr := MajoritySWF(emptyProfile)
	assert.Error(emptyErr) // TODO discuss empty fail

	// Test case equality
	equalityProfile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 1},
	}
	expectedEqualityCount := Count{
		0: 1,
		1: 1,
		2: 1,
	}

	equalityCount, equalityErr := MajoritySWF(equalityProfile)
	assert.NoError(equalityErr)
	assert.DeepEqual(equalityCount, expectedEqualityCount)
}

func TestMajoritySCF(t *testing.T) {
	assert := Assert{t}

	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{0}
	bestAlts, err := MajoritySCF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, expectedBestAlts)
}
