package docking

import (
	"strings"
)

// maxUInt36 is the highest possible Value
// it was originally called maxMask but as exotic as it may be to pronounce, it didn't make much sense
const maxUInt36 = 1<<36 - 1 // 36 ones

// onesReplacer changes all Xes to ones
var onesReplacer = strings.NewReplacer("X", "1")

// zerosReplacer changes all Xes to 0s
var zerosReplacer = strings.NewReplacer("X", "0")

// zerosReplacer changes all non-Xes to 0s and Xes to 1s
var floaterReplacer = strings.NewReplacer("1", "0", "X", "1")
