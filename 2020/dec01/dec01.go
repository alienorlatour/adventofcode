package dec01

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
	"github.com/ablqk/adventofcode/pkg/match"
)

const expectedTotal = 2020

// New instance of the door for December 1st
// 1081 – Birth of Louis VI, French king
func New(inputFile string) doors.Solver {
	return dec01{inputFile}
}

// dec01 is the implementation of this day's exercise
type dec01 struct {
	inputPath string
}

func (d dec01) Solve() (string, error) {
	// read the input file
	lines, err := fileread.ReadInts(d.inputPath)
	if err != nil {
		return "", err
	}
	l, m, err := match.FindSum2(lines, expectedTotal)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d * %d = %d", l, m, l*m), nil
}
