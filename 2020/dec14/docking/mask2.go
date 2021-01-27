package docking

import (
	"fmt"
	"strconv"
)

type Mask2 struct {
	ones         Value // holds 1 where the mask has a one
	exes         Value // holds all possible floatings: a 1 for each X
	floaterCount uint
}

// NewMask parses a mask and returns it
func NewMask2(s string) (Mask2, error) {
	ones, err := strconv.ParseUint(zerosReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask2{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	x, err := strconv.ParseUint(floaterReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask2{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	return Mask2{
		ones: Value(ones),
		exes: Value(x),
		floaterCount: countOnes(x),
	}, nil
}

func countOnes(x uint64) uint {
	var c uint
	for x > 0 {
		if x&1 == 1 {
			c++
		}
		x >>= 1
	}
	return c
}

func (m Mask2) FloaterCount() uint {
	return m.floaterCount
}

// Floater see readme for documentation
func (m Mask2) Floater(v Value, i uint) Value {
	x := m.exes
	// ones will hold bits where we need ones, zeroes will hold the bits where we need zeros
	var ones, zeros Value
	for b := 0; b < 36; b++ { // 36 is too high - this could be optimised for values such as 000000000000000XX0
		if x&1 == 1 { // x ends with 1
			if i&1 == 1 { // i also ends with 1
				ones |= 1 << 36 // we write a one
			} else {
				zeros |= 1 << 36 // we write a one where we will want to force a zero
			}
			// we will next look at the bit to the left
			i >>= 1
		}
		// we will next look at the bit to the left
		ones >>= 1
		zeros >>= 1
		x >>= 1
	}
	// zeros contains ones where we want to force a zero, let's turn it over
	zeros^=maxUInt36

	return v&zeros | ones | m.ones
}
