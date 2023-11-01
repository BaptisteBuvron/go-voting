package comsoc

import (
	"fmt"
	"reflect"
	"testing"
)

type MockProfile struct{}

func (p MockProfile) Voters() []string {
	return []string{"Alice", "Bob", "Charlie"}
}

type MockCount Count

func (c MockCount) Get(alt Alternative) int {
	return c[alt]
}

func mockSWF(p Profile) (Count, error) {
	count := MockCount{
		Alternative(0): 3,
		Alternative(1): 3,
		Alternative(2): 3,
	}
	return Count(count), nil
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
