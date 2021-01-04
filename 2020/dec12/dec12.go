package dec12

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 12
func New(input string) doors.Solver {
	return dec12{input}
}

type dec12 struct {
	input string
}

// Solve the day's problem
func (d dec12) Solve() (string, error) {
	var count int
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}
