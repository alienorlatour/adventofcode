package dec02

import (
	"github.com/ablqk/adventofcode/output"
)

func New() output.Outputter {
	return dec02{}
}

// dec02 is the implementation of this day's exercise
type dec02 struct {}

func (dec02) Output() (string, error) {
	return "", nil
}
