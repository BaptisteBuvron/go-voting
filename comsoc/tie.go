package comsoc

import (
	"errors"
	"sort"
)

func TieBreak(alts []Alternative) (alt Alternative, err error) {
	if len(alts) == 0 {
		err = errors.New("Empty alternatives")
		return
	}
	alt = alts[0]
	return
}

// TODO, le premier est préféré ou destesté ?
func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	return func(alts []Alternative) (Alternative, error) {
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

// les SWF doivent renvoyer un ordre total sans égalité
// TODO TEST
func SWFFactory(swf func(p Profile) (Count, error), tb func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
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
		for alt, _ := range count {
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
func SCFFactory(scf func(p Profile) ([]Alternative, error), tb func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	return func(p Profile) (Alternative, error) {
		bestAlts, err := scf(p)
		if err != nil {
			return -1, err
		}
		if len(bestAlts) == 1 {
			return bestAlts[0], nil
		}
		return tb(bestAlts)
	}
}
