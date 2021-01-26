package docking

import (
	"fmt"
	"strconv"
	"strings"
)

type Mask struct {
	ones  value // holds 1 where the mask has a one
	zeros value // holds 0 where the mask has a zero
}

const maxMask = 1<<36 - 1 // 36 ones

// onesReplacer changes all Xes to ones
var onesReplacer = strings.NewReplacer("X", "1")

// zerosReplacer changes all Xes to 0s
var zerosReplacer = strings.NewReplacer("X", "0")

// NewMask parses a mask and returns it
func NewMask(s string) (Mask, error) {
	// replace all Xes and 0s by 1s
	zeros, err := strconv.ParseUint(onesReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	ones, err := strconv.ParseUint(zerosReplacer.Replace(s), 2, 64)
	if err != nil {
		return Mask{}, fmt.Errorf("invalid mask: %s, %v", s, err)
	}
	return Mask{
		ones:  value(ones),
		zeros: value(zeros),
	}, nil
}
