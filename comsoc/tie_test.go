package comsoc

import (
	"testing"
)

func mockSWF(p Profile) (Count, error) {
	count := Count{
		Alternative(1): 3,
		Alternative(2): 3,
		Alternative(3): 3,
	}
	return count, nil
}

func TestSWFFactory(t *testing.T) {
	assert := Assert{t}

	// Create swf
	tb := TieBreakFactory([]Alternative{3, 1, 2})
	swf := SWFFactory(mockSWF, tb)

	// Act
	got, err := swf(Profile{}) // dummy profile

	// Check result
	assert.NoError(err)
	assert.DeepEqual(got, []Alternative{3, 1, 2})

	// Impossible with the current interface
	// Check with empty profile
	//swf = SWFFactory(MajoritySWF, tb)
	//got, err = swf(Profile{}) // fake profile

	// Check result
	//assert.NoError(err)
	//assert.DeepEqual(got, []Alternative{3, 1, 2})

}

func mockSCF(p Profile) ([]Alternative, error) {
	// For simplicity, returning a fixed set of alternatives
	return []Alternative{1, 2, 3}, nil
}

func TestSCFFactory(t *testing.T) {
	assert := Assert{t}

	// Create scf
	tb := TieBreakFactory([]Alternative{3, 1, 2})
	scf := SCFFactory(mockSCF, tb)
	// Act
	got, err := scf(Profile{}) // dummy profile

	// Check result
	assert.NoError(err)
	assert.DeepEqual(got, Alternative(3))
}
