package dec01

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/fileread"
	"github.com/ablqk/adventofcode/doors"
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
	l, m, err := findMatch2(lines)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d * %d = %d", l, m, l*m), nil
}

func findMatch2(lines []int) (int, int, error) {
	// loop to find 2 values that add up to expectedTotal
	for i, l := range lines[:len(lines)-1] {
		for _, m := range lines[i+1:] {
			if l+m == expectedTotal {
				return l, m, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no match found")
}

