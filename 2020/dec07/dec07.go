package dec07

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 7
func New(input, myBag string) doors.Solver {
	return dec07{input, Colour(myBag)}
}

type dec07 struct {
	input string
	myBag Colour
}

// Solve the day's problem
func (d dec07) Solve() (string, error) {
	bgt := BagRuleTree{}
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return bgt.parseLine(s)
	})
	if err != nil {
		return "", err
	}

	count := len(bgt.Containers(d.myBag).Slice())

	return fmt.Sprintf("%s bags can be found in %d different other-coloured bags", d.myBag, count), nil
}
