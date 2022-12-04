package dec01

import (
	"fmt"
	"strconv"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 25
func New(input string) doors.Solver {
	return dec25{input}
}

type dec25 struct {
	input string
}

// Solve the day's problem
func (d dec25) Solve() (string, error) {
	var count, max int

	err := fileread.ReadAndApply(d.input, func(s string) error {
		if len(s) == 0 {
			// nextElf
			if count > max {
				max = count
			}
			count = 0
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			count += i
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", max), nil
}
