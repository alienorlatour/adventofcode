package dec05

import (
	"fmt"
	"math"
	"strings"

	"github.com/ablqk/adventofcode/2020/fileread"
	"github.com/ablqk/adventofcode/doors"
)

// New instance of the Door for December 5
// 1935 – Mary McLeod Bethune founds the National Council of Negro Women in New York City.
func New(input string) doors.Solver {
	return dec05{input}
}

type dec05 struct {
	input string
}

// Solve the day's problem
func (d dec05) Solve() (string, error) {
	// we need to keep the highest and lowest seats
	var highest BoardingPass
	lowest := BoardingPass{math.MaxInt16, math.MaxInt16} // int16 should leave ample room
	// we also need the sum of all seat IDs and the total number of boarding passes we scanned
	var sum int64 // let's be generous
	var length int
	err := fileread.ReadAndApply(d.input, func(s string) error {
		bp, err := NewBoardingPass(strings.Trim(s, " \t"))
		if err != nil {
			return err
		}
		if bp.IsHigher(highest) {
			highest = bp
		}
		if lowest.IsHigher(bp) {
			lowest = bp
		}
		sum += int64(bp.SeatID())
		length++
		return nil
	})
	if err != nil {
		return "", err
	}
	length++
	// there is probably a problem of luck depending on whether the missing seat is after or before the middle of the plane
	// we compute the expected sum, were all seats actually figuring in the boarding passes
	expectedSum := int64((highest.SeatID()+lowest.SeatID())*length) / 2
	// then we substract the actual sum and return the difference
	return fmt.Sprintf("The highest seat ID is %d", expectedSum-sum), nil
}
