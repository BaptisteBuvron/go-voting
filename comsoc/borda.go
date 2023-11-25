package comsoc

// special case of a scoring method: s₀=0, s₁=1, ..., sₘ−1 = m-1
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#33
var BordaSWF = ScoringSWFFactory(func(index int, size int) int {
	return size - index - 1
})

// See: [comsoc.BordaSWF]
var BordaSCF = SWF2SCF(BordaSWF)
