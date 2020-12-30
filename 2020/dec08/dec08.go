package dec08

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec08/linter"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 8
func New(input string) doors.Solver {
	return dec08{input}
}

type dec08 struct {
	input string
}

// Solve the day's problem
func (d dec08) Solve() (string, error) {
	var comp linter.Compiler
	err := fileread.ReadAndApply(d.input, comp.CompileLine)
	if err != nil {
		return "", err
	}

	// c, _ := comp.Lint()
	c := comp.Patch()

	return fmt.Sprintf("Looping at accumulator value %d", c), nil
}
