package comsoc

import (
	"fmt"
	"reflect"
	"testing"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// Returns the index where alt is located in prefs.
func Rank(alt Alternative, prefs []Alternative) int {
	for index, pref := range prefs {
		if alt == pref {
			return index
		}
	}
	return -1
}

// Returns true iff alt1 is preferred over alt2.
func IsPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return Rank(alt1, prefs) < Rank(alt2, prefs)
}

// Returns the best alternatives for a given count.
func MaxCount(count Count) (bestAlts []Alternative) {
	max := 0
	for alt, total := range count {
		if total > max {
			bestAlts = []Alternative{alt}
			max = total
		} else if total == max {
			bestAlts = append(bestAlts, alt)
		}
	}
	return bestAlts
}

// Checks that the profile is complete and that each alternative only appears once per preference.
func CheckProfile(prefs Profile) error {
	var size int = -1
	if len(prefs) == 0 {
		return fmt.Errorf("Empty Profile")
	}

	for i, alts := range prefs {
		// Check if each preferences have the same size
		if size == -1 {
			size = len(alts)
		} else if size != len(alts) {
			return fmt.Errorf("Incomplete alternatives at %d got %v but expect size %d", i, alts, size)
		}

		// Verify if not alternative are duplicated
		meetAlts := make(map[Alternative]bool)
		for _, alt := range alts {
			if meetAlts[alt] {
				return fmt.Errorf("Duplicate alternative at %d : %d", i, alt)
			}
			meetAlts[alt] = true
		}

		// Verify that all alternatives are present (depending on the size of the profile)
		for i := 0; i < size; i++ {
			alt := Alternative(i)
			if !meetAlts[alt] {
				return fmt.Errorf("Missing alternative at %d : %d", i, alt)
			}
		}
	}
	// No error
	return nil
}

// Checks the profile given, e.g. that they are all complete and that each alts alternative appears exactly once per preferences
func CheckProfileAlternative(prefs Profile, alts []Alternative) error {
	for voter, pref := range prefs {
		missing := make(map[Alternative]bool, len(alts))
		for _, alt := range alts {
			missing[alt] = true
		}
		for index, alt := range pref {
			if missing[alt] {
				missing[alt] = false
			} else {
				return fmt.Errorf("Duplicate alternative %d on voter %d at index %d", alt, voter, index)
			}
		}
		for alt, isMissing := range missing {
			if isMissing {
				return fmt.Errorf("Missing alternative %d on voter %d", alt, voter)
			}
		}
	}
	return nil
}

// Check if alt1 win against alt2 in the given profile
func WinAgainst(alt1 Alternative, alt2 Alternative, p Profile) bool {
	alt1Score := 0
	for _, prefs := range p {
		if IsPref(alt1, alt2, prefs) {
			alt1Score++
		} else {
			alt1Score--
		}
	}
	return 0 < alt1Score
}

// Social Welfare Function.
type SWF func(p Profile) (count Count, err error)

// Social Choice Function.
type SCF func(p Profile) (bestAlts []Alternative, err error)

// Transform a SWF into SCF.
func SWF2SCF(swf SWF) SCF {
	return func(p Profile) (bestAlts []Alternative, err error) {
		count, err := swf(p)
		if err != nil {
			return
		}
		bestAlts = MaxCount(count)
		return
	}
}

func CountFor(p Profile) Count {
	count := make(Count)
	if len(p) == 0 {
		return count
	}
	for _, alt := range p[0] {
		count[alt] = 0
	}
	return count
}

// Function for evaluate score
type ScoreEvaluator func(index int, size int) (score int)

// A voting method that gives points to the candidate based on their position in the ranking
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#33
func ScoringSWFFactory(evaluator ScoreEvaluator) SWF {
	return func(p Profile) (count Count, err error) {
		count = make(Count, 0)
		err = CheckProfile(p)
		if err != nil {
			return
		}
		for _, alts := range p {
			for index, alt := range alts {
				count[alt] += evaluator(index, len(alts))
			}
		}
		return
	}
}

type Assert struct {
	t *testing.T
}

func (a *Assert) NoError(err error) {
	a.t.Helper() // increase stack pointer in log
	if err != nil {
		a.t.Errorf("Unexpected error: %v", err)
	}
}

func (a *Assert) Error(err error) {
	a.t.Helper() // increase stack pointer in log
	if err == nil {
		a.t.Error("An error was expected")
	}
}

func (a *Assert) DeepEqual(got any, expected any) {
	a.t.Helper() // increase stack pointer in log
	if !reflect.DeepEqual(got, expected) {
		a.t.Errorf("Results mismatch: Got %v, Expected %v", got, expected)
	}
}

func (a *Assert) Empty(got any) {
	a.t.Helper()   // increase stack pointer in log
	defer func() { // recover from failed len()
		if r := recover(); r != nil {
			a.t.Errorf("Expected empty array: Got %v", got)
		}
	}()
	if reflect.ValueOf(got).Len() != 0 {
		a.t.Errorf("Expected empty array: Got %v", got)
	}
}

func Alts(alts ...int) []Alternative {
	var alternatives []Alternative = make([]Alternative, len(alts))
	for i, alt := range alts {
		alternatives[i] = Alternative(alt)
	}
	return alternatives
}
