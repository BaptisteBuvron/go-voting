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
		// Classer les counts en fonction du tie break
		//TODO TEST
		sort.SliceStable(count, func(i, j int) bool {
			alt1 := count[Alternative(i)]
			alt2 := count[Alternative(j)]
			if alt1 == alt2 {
				// A partir du tie break on classe le premier et le second
				res, err := TieBreak([]Alternative{Alternative(i), Alternative(j)})
				if err != nil {
					panic(err)
				}
				// On renvoie vrai si le premier est préféré
				return res == Alternative(i)
			}
			return alt1 > alt2
		})
		// Renvoyer les alternatives dans l'ordre
		var alts []Alternative
		for _, c := range count {
			alts = append(alts, Alternative(c))
		}
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
