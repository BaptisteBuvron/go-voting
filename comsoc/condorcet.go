package comsoc

/*
*

	fonction permettant de trouver le gagnant de Condorcet

Qui renvoie un slice éventuellement vide ou ne contenant qu'un seul élément.
*/
func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	var m map[int]int
	//on vérifie que le profil est complet
	err = checkProfile(p)
	if err != nil {
		return bestAlts, err
	}
	//on initialise le tableau de comptage
	m = make(map[int]int)
	//on teste toutes les alternatives
	for i := 0; i < len(p[0]); i++ {
		m[i] = 0
		for j := 0; j < len(p[0]); j++ {
			if i != j {
				for k := 0; k < len(p); k++ {
					if isPref(Alternative(i), Alternative(j), p[k]) {
						m[i]++
					}
				}
			}
		}
	}
	//on cherche le maximum
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	//on cherche les alternatives qui ont le maximum
	for k, v := range m {
		if v == max {
			bestAlts = append(bestAlts, Alternative(k))
		}
	}
	// si on a trouvé une seule alternative, on la renvoie
	if len(bestAlts) == 1 {
		return bestAlts, nil
	}
	//sinon on renvoie un slice vide
	return nil, nil
}
