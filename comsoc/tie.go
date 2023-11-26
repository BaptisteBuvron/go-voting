package comsoc

import (
	"errors"
	"net/http"
	"sort"
)

// A function that allows you to select a candidate among the best.
type TieBreak func([]Alternative) (Alternative, error)

// Take the first candidate.
func TieBreakFirstOne(alts []Alternative) (Alternative, error) {
	if len(alts) == 0 {
		return Alternative(-1), HTTPErrorf(http.StatusBadRequest, "Empty alternatives")
	}
	return alts[0], nil
}

// Take the highest candidate (for be more determinist).
func TieBreakHighest(alts []Alternative) (Alternative, error) {
	if len(alts) == 0 {
		return Alternative(-1), HTTPErrorf(http.StatusBadRequest, "Empty alternatives")
	}
	maxAlt := alts[0]
	for _, alt := range alts {
		if maxAlt < alt {
			maxAlt = alt
		}
	}
	return maxAlt, nil
}

// We prefer the first candidate
func TieBreakFactory(orderedAlts []Alternative) TieBreak {
	return func(alts []Alternative) (Alternative, error) {
		// Check if at least one candidate
		if len(orderedAlts) == 0 {
			return -1, errors.New("Empty alternatives")
		}
		maxIndex := len(orderedAlts)
		maxAlt := alts[0]
		for _, alternative := range alts {
			i := 0
			for i < maxIndex {
				if orderedAlts[i] == alternative {
					maxIndex = i
					maxAlt = alternative
				}
				i += 1
			}
		}
		return maxAlt, nil
	}
}

// Social Welfare Function with tie break.
type SWFWithTieBreak func(Profile) ([]Alternative, error)

// Social Choice Function with tie break.
type SCFWithTieBreak func(Profile) (Alternative, error)

// SWFs must return a total order without ties
func SWFFactory(swf SWF, tb TieBreak) SWFWithTieBreak {
	return func(p Profile) ([]Alternative, error) {
		// Check profile (must be already did by the swf)
		err := CheckProfile(p)
		if err != nil {
			return nil, err
		}
		count, err := swf(p)
		if err != nil {
			return nil, err
		}
		// Rank the counts according to the tie break in case of a tie, otherwise you can large to smallest
		var alts []Alternative
		for alt := range count {
			alts = append(alts, alt)
		}

		// Use the tie break to order the alternatives in case of a tie
		sort.Slice(alts, func(i, j int) bool {
			// Use tie break only in case of tie
			if count[alts[i]] == count[alts[j]] {
				tbResult, _ := tb([]Alternative{alts[i], alts[j]})
				return tbResult == alts[i] // Inverser l'ordre ici
			}
			// If no tie, order by number of votes
			return count[alts[i]] > count[alts[j]]
		})

		return alts, nil
	}
}

// Must return the best alternative regarding to the tiebreak
func SCFFactory(scf SCF, tb TieBreak) SCFWithTieBreak {
	return func(p Profile) (Alternative, error) {
		bestAlts, err := scf(p)
		if err != nil {
			return -1, err
		}
		bestAlt, err := tb(bestAlts)
		if err != nil { // If error, bestAlt must be -1
			return -1, err
		}
		return bestAlt, err
	}
}
