package match

import (
	"fmt"
)

// FindSum2old finds 2 values in an array that add up to expectedTotal
// Deprecated, but I'll keep it to always remember that for most n = len(something), n² > n
func FindSum2old(lines []int, expectedTotal int) (int, int, error) {
	for i, l := range lines[:len(lines)-1] {
		for _, m := range lines[i+1:] {
			if l+m == expectedTotal {
				return l, m, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no match found")
}

// FindSum2 finds 2 values in an array that add up to expectedTotal
func FindSum2(haystack []int, expectedTotal int) (int, int, error) {
	// store complements
	comp := make(map[int]struct{})
	for _, l := range haystack {
		_, found := comp[l]
		if found {
			return expectedTotal-l, l, nil
		}
		comp[expectedTotal-l] = struct{}{}
	}
	return 0, 0, fmt.Errorf("no match found")
}
