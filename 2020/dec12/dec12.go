package dec12

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec12/ship/compiler"
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
	var c compiler.Compiler
	err := fileread.ReadAndApply(d.input, c.ParseLine)
	if err != nil {
		return "", err
	}

	r := c.Compile()
	err = r.Run()
	if err != nil {
		return "Error while running the code: " + err.Error(), nil
	}

	manhattan := abs(r.Longitude()) + abs(r.Latitude())
	return fmt.Sprintf("Manhattan distance is %d.", manhattan), nil
}

// abs returns the absolute value of an int
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
