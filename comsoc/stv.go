package comsoc

// Single Transferable Vote (STV) - Instant Runoff - alternative vote
// ref: https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/td3/sujet.md#rappel-vote-simple-transf%C3%A9rable-single-transferable-vote-stv
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#49
func STV_SWF_Factory(negativeTieBreak TieBreak) SWF {
	return GuardProfile(func(p Profile) (Count, error) {
		// Counter for score
		count := CountFor(p)
		// A set for removed alternatives
		removed := make(map[Alternative]bool, len(p[0]))
		// Iterate over round
		for i := 0; i < len(p[0]); i++ {
			roundCount := CountFor(p)
			// Score At this round
			for _, alts := range p {
				// Find the voted candidate without these removed
				for _, alt := range alts {
					if !removed[alt] {
						roundCount[alt]++
						break
					}
				}
			}

			// Find the lower score or skip if majority
			minScore := 0
			majority := (len(p) - len(removed)) / 2
			for alt, score := range roundCount {

				// We have absolute majority
				if score > majority {
					count[alt]++
					goto STV_SWF_FactoryEnd
				}

				if minScore < score {
					minScore = score
				}
			}

			// Find all losers https://youtu.be/YgSPaXgAdzE
			var losers []Alternative
			for alt, score := range roundCount {
				if score == minScore {
					losers = append(losers, alt)
				}
			}

			// We use TieBreak for choose the biggest loser
			loser, err := negativeTieBreak(losers)
			if err != nil {
				return nil, err
			}
			removed[loser] = true

			// Give points to the winners
			for _, alt := range p[0] {
				if !removed[alt] {
					count[alt]++
				}
			}
		}
	STV_SWF_FactoryEnd:
		return count, nil
	})
}

// See: [comsoc.STV_SWF_Factory]
var STV_SWF = STV_SWF_Factory(TieBreakHighest)

// See: [comsoc.STV_SWF_Factory]
var STV_SCF = SWF2SCF(STV_SWF)
