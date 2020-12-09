package dec07

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/fileread"
	"github.com/ablqk/adventofcode/doors"
)

// New instance of the Door for December 7
func New(input string) doors.Solver {
	return dec07{input}
}

type dec07 struct {
	input string
}

// Solve the day's problem
func (d dec07) Solve() (string, error) {
	var count int
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}
