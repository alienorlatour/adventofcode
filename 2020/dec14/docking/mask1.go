package docking

import (
	"fmt"
	"strconv"
)

type Mask1 struct {
	ones  Value // holds 1 where the mask has a one
	zeros Value // holds 0 where the mask has a zero
}

// NewMask parses a mask and returns it
func NewMask(s string) (Mask1, error) {
	// replace all Xes and 0s by 1s
	zeros, err := strconv.ParseUint(onesReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask1{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	ones, err := strconv.ParseUint(zerosReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask1{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	return Mask1{
		ones:  Value(ones),
		zeros: Value(zeros),
	}, nil
}

func (m Mask1) ApplyTo(v Value) Value {
	return v & m.zeros | m.ones
}
