package comsoc

// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#43
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#43
func CopelandSWF(p Profile) (Count, error) {
	err := CheckProfile(p)
	if err != nil {
		return nil, err
	}
	count := CountFor(p)
	// TODO
	for i, alt := range p[0] {
		for j := i + 1; j < len(p[0]); j++ {
			opponent := p[0][j]
			if WinAgainst(alt, opponent, p) {
				count[alt]++
			} else if WinAgainst(alt, opponent, p) {
				count[alt]--
			}
		}
	}
	return count, nil
}

var CopelandSCF = SWF2SCF(CopelandSWF)
