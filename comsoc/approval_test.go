package comsoc

import (
	"testing"
)

func TestApprovalSWF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 0, 2},
		[]Alternative{2, 1, 0},
	}
	thresholds1 := []int{2, 2, 2}

	expectedCout := Count{
		0: 2,
		1: 3,
		2: 1,
	}

	count, err := ApprovalSWF(profile1, thresholds1)
	assert.NoError(err)
	assert.DeepEqual(count, expectedCout)

	thresholds2 := []int{3, 2, 2}

	expectedCout2 := Count{
		0: 2,
		1: 3,
		2: 2,
	}

	count2, err := ApprovalSWF(profile1, thresholds2)
	assert.NoError(err)
	assert.DeepEqual(count2, expectedCout2)

	thresholds3 := []int{3, 3, 3}
	expectedCount3 := Count{
		0: 3,
		1: 3,
		2: 3,
	}

	count3, err := ApprovalSWF(profile1, thresholds3)
	assert.NoError(err)
	assert.DeepEqual(count3, expectedCount3)

}

func TestApprovalSCF(t *testing.T) {
	assert := Assert{t}

	profile1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 0, 2},
		[]Alternative{2, 1, 0},
	}
	thresholds1 := []int{2, 2, 2}

	bestAlts, err := ApprovalSCF(profile1, thresholds1)
	assert.NoError(err)
	assert.DeepEqual(bestAlts, []Alternative{1})

	thresholds2 := []int{1, 1, 1}
	expectedBestAlts2 := []Alternative{0, 1, 2}
	bestAlts2, err := ApprovalSCF(profile1, thresholds2)
	assert.NoError(err)
	assert.DeepEqual(bestAlts2, expectedBestAlts2)

	thresholds3 := []int{3, 3, 2}
	expectedBestAlts3 := []Alternative{1, 2}
	bestAlts3, err := ApprovalSCF(profile1, thresholds3)
	assert.NoError(err)
	assert.DeepEqual(bestAlts3, expectedBestAlts3)

}
