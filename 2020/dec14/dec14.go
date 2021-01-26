package dec14

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec14/docking"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 14
func New(input string) doors.Solver {
	return dec14{input}
}

type dec14 struct {
	input string
}

// Solve the day's problem
func (d dec14) Solve() (string, error) {
	p := Parser{
		memory: make(map[string]docking.Value),
	}
	err := fileread.ReadAndApply(d.input, p.Parse)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Final addition is %d.", p.Count()), nil
}
