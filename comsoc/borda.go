package comsoc

func BordaSWF(p Profile) (count Count, err error) {
	err = CheckProfile(p)
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

var BordaSCF = SWF2SCF(BordaSWF)
