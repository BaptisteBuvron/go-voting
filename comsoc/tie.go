package comsoc

import (
	"errors"
	"sort"
)

// A function that allows you to select a candidate among the best.
type TieBreak func(alts []Alternative) (alt Alternative, err error)

// Take the first candidate.
func TieBreakFirstOne(alts []Alternative) (Alternative, error) {
	if len(alts) == 0 {
		return Alternative(-1), errors.New("Empty alternatives")
	}
	return alts[0], nil
}

// Take the highest candidate (for be more determinist).
func TieBreakHighest(alts []Alternative) (alt Alternative, err error) {
	if len(alts) == 0 {
		return Alternative(-1), errors.New("Empty alternatives")
	}
	maxAlt := alts[0]
	for _, alt := range alts {
		if maxAlt < alt {
			maxAlt = alt
		}
	}
	return maxAlt, nil
}

// TODO, le premier est préféré ou détesté ?
func TieBreakFactory(orderedAlts []Alternative) TieBreak {
	// TODO make verification
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
type SWFWithTieBreak func(p Profile) (alts []Alternative, err error)

// Social Choice Function with tie break.
type SCFWithTieBreak func(p Profile) (bestAlt Alternative, err error)

// les SWF doivent renvoyer un ordre total sans égalité
// TODO TEST
func SWFFactory(swf SWF, tb TieBreak) SWFWithTieBreak {
	return func(p Profile) ([]Alternative, error) {
		err := CheckProfile(p)
		if err != nil {
			return nil, err
		}
		count, err := swf(p)
		if err != nil {
			return nil, err
		}
		// Classer les counts en fonction du tie break en cas d'égalité, sinon tu peux grand au plus petit
		var alts []Alternative
		for alt := range count {
			alts = append(alts, alt)
		}

		// Utiliser le tie break pour ordonner les alternatives en cas d'égalité
		sort.Slice(alts, func(i, j int) bool {
			// Utiliser le tie break uniquement en cas d'égalité
			if count[alts[i]] == count[alts[j]] {
				tbResult, _ := tb([]Alternative{alts[i], alts[j]})
				return tbResult == alts[i] // Inverser l'ordre ici
			}
			// Si pas d'égalité, ordonner par le nombre de voix
			return count[alts[i]] > count[alts[j]]
		})

		return alts, nil
	}
}

// TODO TEST
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
