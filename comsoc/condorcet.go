package comsoc

// checks if a Condorcet winner exists, if this is not the case return an empty slice
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#13
func CondorcetWinner(p Profile) ([]Alternative, error) {
	// Empty case
	if len(p) == 0 {
		return nil, nil
	}
	// Check if profile is complet
	err := CheckProfile(p)
	if err != nil {
		return nil, err
	}
	// Tests on all alternatives
	var bestAlts []Alternative
	for _, alt1 := range p[0] {
		// Search if this alternative fail against someone
		for _, alt2 := range p[0] {
			if alt1 != alt2 {
				if !WinAgainst(alt1, alt2, p) {
					goto failed // we can stop here, he lose against someone
				}
			}
		}
		return append(bestAlts, alt1), nil
	failed: // Lose against someone
	}
	// Return empty slice with no errors
	return nil, nil
}
