package dec17

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec17/energysource/source4d"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 17
func New(input string, iterations int) Dec17 {
	return Dec17{
		input:      input,
		iterations: iterations,
	}
}

type Dec17 struct {
	input      string
	iterations int
}

// Solve the day's problem
func (d Dec17) Solve() (string, error) {
	state, err := newState(d.input)
	if err != nil {
		return "", fmt.Errorf("cannot parse input file: %w", err)
	}

	actives := state.CountActiveCellsAfter(d.iterations)
	return fmt.Sprintf("After %d iterations, we have %d active cells", d.iterations, actives), nil
}

// Source of Experimental Energy
type source interface {
	CountActiveCellsAfter(times int) int
}

// newState parses the input file and returns a source
func newState(input string) (source, error) {
	state := source4d.New()
	var y int
	err := fileread.ReadAndApply(input, func(s string) error {
		for x, c := range s {
			if c == '#' {
				state.Activate(x, y, 0, 0)
			}
		}
		y++
		return nil
	})
	return state, err
}
