package dec05

import (
	"github.com/ablqk/adventofcode/doors"
)

func New(input string) doors.Solver {
	return dec05{input}
}

type dec05 struct {
	input string
}

func (d dec05) Solve() (string, error) {
	return "", nil
}
