package comsoc

import (
	"errors"
)

func TieBreak(alts []Alternative) (alt Alternative, err error) {
	if len(alts) == 0 {
		err = errors.New("Empty alternatives")
		return
	}
	alt = alts[0]
	return
}
