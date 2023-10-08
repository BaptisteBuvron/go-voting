package comsoc

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

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	if err != nil {
		return
	}
	bestAlts = MaxCount(count)
	return
}
