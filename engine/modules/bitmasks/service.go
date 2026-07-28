package bitmasks

import (
	"math/bits"

	"golang.org/x/exp/constraints"
)

func GetFlags[BitMask constraints.Unsigned](f BitMask) []BitMask {
	var flags []BitMask
	flags = make([]BitMask, 0, bits.OnesCount64(uint64(f)))

	var flag BitMask = 1
	for f != 0 {
		if f&flag != 0 {
			flags = append(flags, flag)
			f &= ^flag
		}
		flag <<= 1
	}
	return flags
}
