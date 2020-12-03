package dec04

import (
	"github.com/ablqk/adventofcode/doors"
)

func New(input string) doors.Solver {
	return dec04{input}
}

type dec04 struct {
	input string
}

func (d dec04) Solve() (string, error) {
	return "", nil
}
