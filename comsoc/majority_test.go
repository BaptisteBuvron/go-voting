package comsoc

import (
	"reflect"
	"testing"
)

func TestMajoritySWF(t *testing.T) {
	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}
	expectedCount := Count{
		0: 2,
		1: 1,
	}

	count, err := MajoritySWF(validProfile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(count, expectedCount) {
		t.Errorf("Expected count %v, but got %v", expectedCount, count)
	}

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{0}
	bestAlts, err := MajoritySCF(validProfile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(bestAlts, expectedBestAlts) {
		t.Errorf("Expected best alternatives %v, but got %v", expectedBestAlts, bestAlts)
	}

	// Test case  Empty profile
	emptyProfile := Profile{}
	expectedEmptyCount := Count{}

	emptyCount, emptyErr := MajoritySWF(emptyProfile)
	if emptyErr != nil {
		t.Errorf("Unexpected error for empty profile: %v", emptyErr)
	}

	if !reflect.DeepEqual(emptyCount, expectedEmptyCount) {
		t.Errorf("Expected empty count %v, but got %v", expectedEmptyCount, emptyCount)
	}

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
	if equalityErr != nil {
		t.Errorf("Unexpected error for equality profile: %v", equalityErr)
	}

	if !reflect.DeepEqual(equalityCount, expectedEqualityCount) {
		t.Errorf("Expected equality count %v, but got %v", expectedEqualityCount, equalityCount)
	}

}

func TestMajoritySCF(t *testing.T) {
	// Test case 1: Valid profile
	validProfile := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{0, 2, 1},
	}

	// Test case MajoritySCF
	expectedBestAlts := []Alternative{0}
	bestAlts, err := MajoritySCF(validProfile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(bestAlts, expectedBestAlts) {
		t.Errorf("Expected best alternatives %v, but got %v", expectedBestAlts, bestAlts)
	}
}
