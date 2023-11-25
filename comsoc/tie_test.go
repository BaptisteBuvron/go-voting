package comsoc

import (
	"testing"
)

var dummyProfile = Profile{[]Alternative{0}}

func mockSWF(p Profile) (Count, error) {
	count := Count{
		Alternative(0): 3,
		Alternative(1): 3,
		Alternative(2): 3,
	}
	return count, nil
}

func TestSWFFactory(t *testing.T) {
	assert := Assert{t}

	// Create swf
	tb := TieBreakFactory(Alts(2, 0, 1))
	swf := SWFFactory(mockSWF, tb)

	// Act
	got, err := swf(dummyProfile) // fake profile

	// Check result
	assert.NoError(err)
	assert.DeepEqual(got, Alts(2, 0, 1))
}

func mockSCF(p Profile) ([]Alternative, error) {
	// For simplicity, returning a fixed set of alternatives
	return Alts(0, 1, 2), nil
}

func TestSCFFactory(t *testing.T) {
	assert := Assert{t}

	// Create scf
	tb := TieBreakFactory(Alts(2, 0, 1))
	scf := SCFFactory(mockSCF, tb)
	// Act
	got, err := scf(dummyProfile) // fake profile

	// Check result
	assert.NoError(err)
	assert.DeepEqual(got, Alternative(2))
}
