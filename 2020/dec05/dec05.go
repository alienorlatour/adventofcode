package dec05

import (
	"fmt"
	"strings"

	"github.com/ablqk/adventofcode/2020/fileread"
	"github.com/ablqk/adventofcode/doors"
)

func New(input string) doors.Solver {
	return dec05{input}
}

type dec05 struct {
	input string
}

func (d dec05) Solve() (string, error) {
	var highest BoardingPass
	err := fileread.ReadAndApply(d.input, func(s string) error {
		bp, err := NewBoardingPass(strings.Trim(s, " \t"))
		if err != nil {
			return err
		}
		if bp.IsHigher(highest) {
			highest = bp
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("The highest seat ID is %d", highest.SeatID()), nil
}
