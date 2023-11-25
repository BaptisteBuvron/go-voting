package comsoc

// A voting method that allows you to vote for as many candidates as you want
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#32
func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = CheckProfile(p)
	if err != nil {
		return
	}
	for voter, alts := range p {
		for i := 0; i < thresholds[voter]; i++ {
			alt := alts[i]
			count[alt] += 1
		}
	}
	return
}

// See: [comsoc.ApprovalSWF]
func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	if err != nil {
		return
	}
	bestAlts = MaxCount(count)
	return
}
