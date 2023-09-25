package comsoc

import (
	"testing"
)

func TestCheckProfile(t *testing.T) {
	var err error
	//Should succeed
	prefs1 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 1},
	}
	err = checkProfile(prefs1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	//Should fail
	prefs2 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 2, 3},
	}
	err = checkProfile(prefs2)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	//Should fail
	prefs3 := Profile{
		[]Alternative{0, 1, 2},
		[]Alternative{1, 2, 0},
		[]Alternative{2, 0, 3},
	}
	err = checkProfile(prefs3)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	//Should fail
	prefs4 := Profile{
		[]Alternative{0, 1, 3},
	}
	err = checkProfile(prefs4)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
}