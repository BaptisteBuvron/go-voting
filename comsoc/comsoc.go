package comsoc

import (
	"errors"
	"fmt"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for index, pref := range prefs {
		if alt == pref {
			return index
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est preferee a alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return rank(alt1, prefs) < rank(alt2, prefs)
}

// renvoie les meilleures alternatives pour un decomtpe donne
func maxCount(count Count) (bestAlts []Alternative) {
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

// verifie le profil donne, par ex. qu'ils sont tous complets et que chaque alternative n'apparait qu'une seule fois par preferences
func checkProfile(prefs Profile) error {
	var size int = -1
	for _, alternative := range prefs {
		alts := make(map[Alternative]int)
		if size == -1 {
			size = len(alternative)
		} else if size != len(alternative) {
			return errors.New("Incomplete profile")
		}
		for _, a := range alternative {
			if alts[a] == 1 {
				return errors.New("Duplicate alternative")
			}
			alts[a] = 1
		}
		//verify that all alternatives are present (depending on the size of the profile)
		for i := 0; i < size; i++ {
			if alts[Alternative(i)] != 1 {
				return errors.New("Incomplete profile")
			}
		}
	}
	return nil
}

// verifie le profil donne, par ex. qu'ils sont tous complets et que chaque alternative de alts apparait exactement une fois par preferences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for voter, pref := range prefs {
		missing := make(map[Alternative]bool, len(alts))
		for _, alt := range alts {
			missing[alt] = true;
		}
		for index, alt := range pref {
			if missing[alt] {
				missing[alt] = false
			} else {
				return fmt.Errorf("Duplicate alternative %d on voter %d at index %d", alt, voter, index)
			}
		}
		for alt, isMissing := range missing {
			if (isMissing) {
				return fmt.Errorf("Missing alternative %d on voter %d", alt, voter)
			}
		}
	}
	return nil
}
