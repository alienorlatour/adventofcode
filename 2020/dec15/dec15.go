package dec15

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
)

// New instance of the Door for December 15
func New(input []int, part1, part2 int) doors.Solver {
	return dec15{input: input, end1: part1, end2: part2}
}

type dec15 struct {
	input []int
	end1  int // if a 3rd part was a possibility, an array would be better
	end2  int
}

// Solve the day's problem
func (d dec15) Solve() (string, error) {
	m := NewMemoryGame(d.input)
	v1 := m.ValueAt(d.end1)
	v2 := m.ValueAt(d.end2)
	return fmt.Sprintf("Value at %d is %d and %d is %d.", d.end1, v1, d.end2, v2), nil
}
