package dec06

import (
	"github.com/ablqk/adventofcode/doors"
)

func New(input string) doors.Solver {
	return dec06{input}
}

type dec06 struct {
	input string
}

func (d dec06) Solve() (string, error) {
	return "", nil
}
