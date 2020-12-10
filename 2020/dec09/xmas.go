package dec09

import (
	"fmt"
)

type XMAS struct {
	preamble int
}

func (x XMAS) FindFirstInvalid(input []int) int {
	i := x.preamble
	// todo check that input length > preamble
	// todo if everything is valid, you will explode
	for ; x.IsValid(input, i); i++ {
	}
	return input[i]
}

func (x XMAS) IsValid(input []int, index int) bool {
	_, _, err := findMatch2(input[index-x.preamble:index], input[index])
	return err == nil       // no match found is the error
}

// findMatch2 is a reuse from December 1st
func findMatch2(lines []int, expectedTotal int) (int, int, error) {
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
