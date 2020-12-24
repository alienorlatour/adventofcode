package match

import (
	"fmt"
)

// FindMatch2 finds 2 values in an array that add up to expectedTotal
func FindMatch2(lines []int, expectedTotal int) (int, int, error) {
	for i, l := range lines[:len(lines)-1] {
		for _, m := range lines[i+1:] {
			if l+m == expectedTotal {
				return l, m, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no match found")
}
