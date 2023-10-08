package comsoc

//TODO TEST
func MajoritySWF(p Profile) (count Count, err error) {
	count = make(Count, 0)
	err = CheckProfile(p)
	if err != nil {
		return
	}
	for _, alts := range p {
		if len(alts) != 0 {
			alt := alts[0]
			count[alt] += 1
		}
	}
	return
}

var MajoritySCF = SWF2SCF(MajoritySWF)
