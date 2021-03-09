package dec17

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec17/energysource"
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

	state = state.Iterate(d.iterations)

	return fmt.Sprintf("Afetr %d iterations, we have %d active cells", d.iterations, state.CountActiveCells()), nil
}

// newState parses the input file
func newState(input string) (*energysource.State, error) {
	state := energysource.New()
	var y int
	err := fileread.ReadAndApply(input, func(s string) error {
		for x, c := range s {
			if c == '#' {
				state.Activate(x, y, 0)
			}
		}
		y++
		return nil
	})
	return state, err
}
