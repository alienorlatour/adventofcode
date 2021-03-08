package dec02

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the door for December 2nd
// 1804 – At Notre Dame Cathedral in Paris, Napoleon Bonaparte crowns himself Emperor of the French.
func New(inputFile string) doors.Solver {
	return dec02{inputFile}
}

// dec02 is the implementation of this day's exercise
type dec02 struct {
	inputPath string
}

func (d dec02) Solve() (string, error) {
	// read the input file
	var valid int
	err := fileread.ReadAndApply(d.inputPath, func(s string) error {
		p, err := newPasswordFromLine(s)
		if err != nil {
			return err
		}
		// if this line is valid, increment the counter
		if p.IsValidNewStyle() {
			valid++
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d valid passwords", valid), nil
}
