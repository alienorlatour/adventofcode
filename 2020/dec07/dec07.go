package dec07

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec07/defs"
	"github.com/ablqk/adventofcode/2020/dec07/part2"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 7
func New(input, myBag string) doors.Solver {
	return dec07{input, defs.Colour(myBag)}
}

type dec07 struct {
	input string
	myBag defs.Colour
}

// Solve the day's problem
func (d dec07) Solve() (string, error) {
	bgt := part2.NewBRT()
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return bgt.ParseLine(s)
	})
	if err != nil {
		return "", err
	}

	count := bgt.Count(d.myBag)

	return fmt.Sprintf("%s bags can be found in %d different other-coloured bags", d.myBag, count), nil
}
