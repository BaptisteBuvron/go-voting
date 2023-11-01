package comsoc

import (
	"fmt"
	"reflect"
	"testing"
)

func mockSWF(p Profile) (Count, error) {
	count := Count{
		Alternative(0): 3,
		Alternative(1): 3,
		Alternative(2): 3,
	}
	return count, nil
}

func mockTieBreak(alt []Alternative) (Alternative, error) {
	return TieBreakFactory([]Alternative{Alternative(2), Alternative(0), Alternative(1)})(alt)
}

func TestSWFFactory(t *testing.T) {
	// Arrange
	factory := SWFFactory(mockSWF, mockTieBreak)
	expectedResult := []Alternative{Alternative(2), Alternative(0), Alternative(1)}

	// Act
	result, err := factory(Profile{})
	fmt.Println(result)

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the result is sorted correctly
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result does not match expected order. Got %v, expected %v", result, expectedResult)
	}
}

func mockSCF(p Profile) ([]Alternative, error) {
	// For simplicity, returning a fixed set of alternatives
	return []Alternative{Alternative(0), Alternative(1), Alternative(2)}, nil
}

func TestSCFFactory(t *testing.T) {
	// Arrange
	factory := SCFFactory(mockSCF, mockTieBreak)
	expectedResult := Alternative(2)

	// Act
	result, err := factory(Profile{})

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the result matches the expected alternative
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result does not match expected alternative. Got %v, expected %v", result, expectedResult)
	}
}
