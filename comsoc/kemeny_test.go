package comsoc

import (
	"testing"
)

func TestKemenySWF(t *testing.T) {
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
	count, err := KemenySWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case 2
	validProfile = Profile{
		[]Alternative{1, 3, 2},
		[]Alternative{2, 3, 1},
		[]Alternative{1, 3, 2},
		[]Alternative{3, 1, 2},
		[]Alternative{2, 3, 1},
	}
	expectedCount = Count{
		1: 1,
		2: 0,
		3: 2,
	}
	count, err = KemenySWF(validProfile)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	// Test case  Empty profile
	_, emptyErr := KemenySWF(Profile{})
	assert.NoError(emptyErr)
}

func TestProfileKendallTauDistance(t *testing.T) {
	assert := NewAssert(t)
	alts := []Alternative{1, 2, 3, 4}
	p1 := Profile{[]Alternative{2, 1, 3, 4}}
	p2 := Profile{[]Alternative{2, 1, 4, 3}}
	p3 := Profile{[]Alternative{4, 2, 3, 1}}
	assert.Equal(ProfileKendallTauDistance(alts, p1), uint64(1))
	assert.Equal(ProfileKendallTauDistance(alts, p2), uint64(2))
	assert.Equal(ProfileKendallTauDistance(alts, p3), uint64(5))
}
