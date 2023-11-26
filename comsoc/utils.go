package comsoc

import (
	"fmt"
	"net/http"
	"reflect"
	"sort"
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

// Get a map for get rank of Alternative
func Ranker(prefs []Alternative) Count {
	ranker := make(Count)
	for index, pref := range prefs {
		ranker[pref] = index
	}
	return ranker
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
	//order bestAlts
	sort.Slice(bestAlts, func(i, j int) bool {
		return bestAlts[i] < bestAlts[j]
	})
	return bestAlts
}

// Check if alternatives are valid
func CheckAlternatives(alts []Alternative, size int) error {
	// Check if each preferences have the same size
	if size != len(alts) {
		return HTTPErrorf(http.StatusBadRequest, "Incomplete alternatives got %v but expect size %d", alts, size)
	}

	// Verify if not alternative are duplicated
	meetAlts := make(map[Alternative]bool)
	for _, alt := range alts {
		if meetAlts[alt] {
			return HTTPErrorf(http.StatusBadRequest, "Duplicate alternative %d", alt)
		}
		meetAlts[alt] = true
	}

	// Verify that all alternatives are present (depending on the size of the profile)
	for i := 1; i <= size; i++ {
		alt := Alternative(i)
		if !meetAlts[alt] {
			return HTTPErrorf(http.StatusBadRequest, "Missing alternative at %d", alt)
		}
	}

	// No error found
	return nil
}

// Checks that the profile is complete and that each alternative only appears once per preference.
func CheckProfile(prefs Profile) error {
	if len(prefs) == 0 {
		return nil // empty profile
	}
	size := len(prefs[0])
	// Check all alternatives
	for _, alts := range prefs {
		err := CheckAlternatives(alts, size)
		if err != nil {
			return err
		}
	}
	// No error
	return nil
}

// Checks that the profile is complete adn skip if empty
func GuardProfile(swf SWF) SWF {
	return func(p Profile) (Count, error) {
		if len(p) == 0 {
			return make(Count), nil
		}
		err := CheckProfile(p)
		if err != nil {
			return nil, err
		}
		return swf(p)
	}
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
				return HTTPErrorf(http.StatusBadRequest, "Duplicate alternative %d on voter %d at index %d", alt, voter, index)
			}
		}
		for alt, isMissing := range missing {
			if isMissing {
				return HTTPErrorf(http.StatusBadRequest, "Missing alternative %d on voter %d", alt, voter)
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
type SWF func(Profile) (Count, error)

// Social Choice Function.
type SCF func(Profile) ([]Alternative, error)

// Transform a SWF into SCF.
func SWF2SCF(swf SWF) SCF {
	return func(p Profile) ([]Alternative, error) {
		count, err := swf(p)
		if err != nil {
			return nil, err
		}
		return MaxCount(count), nil
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
type ScoreEvaluator func(int, int) int

// A voting method that gives points to the candidate based on their position in the ranking
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#33
func ScoringSWFFactory(evaluator ScoreEvaluator) SWF {
	return func(p Profile) (Count, error) {
		count := CountFor(p)
		err := CheckProfile(p)
		if err != nil {
			return nil, err
		}
		for _, alts := range p {
			for index, alt := range alts {
				count[alt] += evaluator(index, len(alts))
			}
		}
		return count, nil
	}
}

type Assert struct {
	t *testing.T
}

func NewAssert(t *testing.T) Assert {
	return Assert{t}
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

func (a *Assert) True(boolean bool) {
	a.t.Helper() // increase stack pointer in log
	if boolean {
		a.t.Error("Boolean must be true")
	}
}

func (a *Assert) DeepEqual(got any, expected any) {
	a.t.Helper() // increase stack pointer in log
	if !reflect.DeepEqual(got, expected) {
		a.t.Errorf("Results mismatch: Got %v, Expected %v", got, expected)
	}
}

func (a *Assert) Equal(got any, expected any) {
	a.t.Helper() // increase stack pointer in log
	if got != expected {
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

type HTTPError struct {
	Code    int
	Message string
}

func (e HTTPError) Error() string {
	return e.Message
}

func (e HTTPError) StatusCode() int {
	return e.Code
}

func HTTPErrorf(status int, message string, args ...any) HTTPError {
	return HTTPError{status, fmt.Sprintf(message, args...)}
}
