package dec09

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 9
func New(input string, preambleSize int) doors.Solver {
	return dec09{input: input, preambleSize: preambleSize}
}

type dec09 struct {
	input        string
	preambleSize int
}

// Solve the day's problem
func (d dec09) Solve() (string, error) {
	input, err := fileread.ReadInts(d.input)
	if err != nil {
		return "", err
	}

	c := XMAS{d.preambleSize}.FindFirstInvalid(input)

	w := findWeakness(input, c)

	return fmt.Sprintf("First invalid value is %d, weakness is %d", c, w), nil
}

const minimalWeaknessSize = 2

// findWeakness for this piece of encryption
func findWeakness(input []int, total int) int {
	lowBound := 0
	highBound := minimalWeaknessSize
	// sum just needs to be different than total
	for sum := addUp(input[lowBound:highBound]); sum != total; {
		highBound++
		if sum > total || highBound >= len(input) {
			// try starting one step further
			lowBound++
			highBound = lowBound + minimalWeaknessSize
		}
		sum = addUp(input[lowBound:highBound])
	}
	small, big := bounds(input[lowBound:highBound])
	return small + big
}

func addUp(ints []int) int {
	var count int
	for _, i := range ints {
		count += i
	}
	return count
}

func bounds(ints []int) (small, big int) {
	// todo check length of ints
	small = ints[0]
	big = ints[0]
	for _, i := range ints {
		if i > big {
			big = i
		}
		if i < small {
			small = i
		}
	}
	return small, big
}
