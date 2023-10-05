package comsoc

//TODO TEST
func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfile(p)
	if err != nil {
		return
	}
	for _, alts := range p {
		if len(alts) != 0 {
			alt := alts[0]
			count[alt] += 1
			break
		}
	}
	return
}

//TODO TEST
func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	count, err := MajoritySWF(p)
	if err != nil {
		return
	}
	bestAlts = maxCount(count)
	return
}
