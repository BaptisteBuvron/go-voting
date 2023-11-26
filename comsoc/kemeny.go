package comsoc

import "math"

// ReverseAlternativeSlice reverses the given slice
// ref: https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/utils/permutations.go#L4
func ReverseAlternativeSlice(a []Alternative) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

// FirstPermutation returns the slice of int [0, 1, ..., n-1]
// ref: https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/utils/permutations.go#L22
func FirstPermutation(n int) (a []Alternative) {
	a = make([]Alternative, n)
	for i := 0; i < n; i++ {
		a[i] = Alternative(i + 1)
	}
	return
}

// NextPermutation returns the next lexicographical permutation of an integer slice
// Perform operation in place
// ref: https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/utils/permutations.go#L32
func NextPermutation(a []Alternative) bool {
	n := len(a)

	if n < 2 {
		return false
	}

	k := n - 2
	for a[k] >= a[k+1] {
		k--
		if k < 0 {
			return false
		}
	}

	l := n - 1
	for a[k] >= a[l] {
		l--
	}

	a[k], a[l] = a[l], a[k]
	ReverseAlternativeSlice(a[k+1:])
	return true
}

// Get the distance of alternatives between all alternatives in profiles
// ref: https://en.wikipedia.org/wiki/Kemeny%E2%80%93Young_method
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#43
func ProfileKendallTauDistance(alts []Alternative, profile Profile) int {
	var distance int = 0
	ranker := Ranker(alts)
	for _, otherAlts := range profile {
		// Kendall Tau distance
		// ref: https://en.wikipedia.org/wiki/Kendall_tau_distance
		// ref: https://jamesmccaffrey.wordpress.com/2021/11/22/the-kendall-tau-distance-for-permutations-example-python-code/
		for i, alt := range otherAlts {
			for _, opponent := range otherAlts[i+1:] {
				if ranker[alt] > ranker[opponent] {
					distance++
				}
			}
		}
	}
	return distance
}

// Find best result by get distance of profile
// ref: https://gitlab.utc.fr/lagruesy/ia04/-/blob/main/docs/sujets/td6/sujet.md
// ref: https://en.wikipedia.org/wiki/Kemeny%E2%80%93Young_method
// ref: https://www.hds.utc.fr/~lagruesy/ens/ia04/02-Prise%20de%20d%C3%A9cision%20collective%20et%20th%C3%A9orie%20du%20choix%20social/#43
var KemenySWF = GuardProfile(func(p Profile) (Count, error) {
	distances := make(map[Alternative]int)
	for _, alt := range p[0] {
		distances[alt] = math.MaxInt // Max Value
	}
	maxDistance := 0
	var ok bool = true
	for alts := FirstPermutation(len(p[0])); ok; ok = NextPermutation(alts) {
		distance := ProfileKendallTauDistance(alts, p)
		alt := alts[0]
		if distances[alt] > distance {
			distances[alt] = distance
		}
		if maxDistance < distance {
			maxDistance = distance
		}
	}
	// revert and normalize count
	for alt, distance := range distances {
		distances[alt] = maxDistance - distance - 1
	}
	return distances, nil
})

// See: [comsoc.KemenySWF]
var KemenySCF = SWF2SCF(KemenySWF)
