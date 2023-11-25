package comsoc

// special case of a scoring method: s₀ = s₁ = ... = sₘ₋₂ < sₘ₋₁
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#33
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%c3%a9cision%20collective%20et%20th%c3%a9orie%20du%20choix%20social/#11
var MajoritySWF = ScoringSWFFactory(func(index int, size int) int {
	if index == 0 {
		return 1
	} else {
		return 0
	}
})

// See: [comsoc.MajoritySWF]
var MajoritySCF = SWF2SCF(MajoritySWF)
