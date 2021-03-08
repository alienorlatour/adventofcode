package dec11

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec11/seating"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 11
func New(input string) doors.Solver {
	return dec11{input}
}

type dec11 struct {
	input string
}

// Solve the day's problem
func (d dec11) Solve() (string, error) {
	var a seating.Arrangement
	err := fileread.ReadAndApply(d.input, a.ParseLine)
	if err != nil {
		return "", err
	}

	l := seating.NewLooper(a)
	people := l.Stabilise()

	return fmt.Sprintf("After stabilisation, there are %d people in the room", people), nil
}
