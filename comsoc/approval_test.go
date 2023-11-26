package comsoc

import (
	"testing"
)

func TestApprovalSWF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 1, 3},
		[]Alternative{3, 2, 1},
	}
	thresholds1 := []int{2, 2, 2}

	expectedCount := Count{
		1: 2,
		2: 3,
		3: 1,
	}

	count, err := ApprovalSWF(profile1, thresholds1)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCount)

	thresholds2 := []int{3, 2, 2}

	expectedCount2 := Count{
		1: 2,
		2: 3,
		3: 2,
	}

	count2, err := ApprovalSWF(profile1, thresholds2)
	assert.NoError(err)
	assert.DeepEqual(count2, expectedCount2)

	thresholds3 := []int{3, 3, 3}
	expectedCount3 := Count{
		1: 3,
		2: 3,
		3: 3,
	}

	count3, err := ApprovalSWF(profile1, thresholds3)
	assert.NoError(err)
	assert.DeepEqual(count3, expectedCount3)

	// Test case  Empty profile
	_, emptyErr := ApprovalSWF(Profile{}, []int{})
	assert.NoError(emptyErr)

	// Test case  Wrong threshold
	_, emptyErr = ApprovalSWF(Profile{}, []int{0})
	assert.Error(emptyErr)
}

func TestApprovalSCF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{1, 2, 3},
		[]Alternative{2, 1, 3},
		[]Alternative{3, 2, 1},
	}
	thresholds1 := []int{2, 2, 2}

	bestAlts, err := ApprovalSCF(profile1, thresholds1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, []Alternative{2})

	thresholds2 := []int{1, 1, 1}
	expectedBestAlts2 := []Alternative{1, 2, 3}
	bestAlts2, err := ApprovalSCF(profile1, thresholds2)
	assert.NoError(err)
	assert.DeepEqual(bestAlts2, expectedBestAlts2)

	thresholds3 := []int{3, 3, 2}
	expectedBestAlts3 := []Alternative{2, 3}
	bestAlts3, err := ApprovalSCF(profile1, thresholds3)
	assert.NoError(err)
	assert.DeepEqual(bestAlts3, expectedBestAlts3)
}
