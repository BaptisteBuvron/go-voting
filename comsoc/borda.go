package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return
	}
	for _, alts := range p {
		for n, alt := range alts {
			count[alt] += len(alts) - n - 1
		}
	}
	return
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := BordaSWF(p)
	if err != nil {
		return
	}
	bestAlts = maxCount(count)
	return
}
