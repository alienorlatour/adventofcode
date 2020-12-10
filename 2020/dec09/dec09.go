package dec09

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 9
func New(input string, preambleSize int) doors.Solver {
	return dec09{input: input, preambleSize: preambleSize}
}

type dec09 struct {
	input        string
	preambleSize int
}

// Solve the day's problem
func (d dec09) Solve() (string, error) {
	input, err := fileread.ReadInts(d.input)
	if err != nil {
		return "", err
	}

	c := XMAS{d.preambleSize}.FindFirstInvalid(input)

	return fmt.Sprintf("First invalid value is %d", c), nil
}
