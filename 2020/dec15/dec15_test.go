package dec15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec15_Solve(t *testing.T) {
	tt := map[string]struct {
		in        []int
		expected1 int
		expected2 int
	}{
		"132": {
			in:        []int{1, 3, 2},
			expected1: 1,
			expected2: 2578,
		},
		"213": {
			in:        []int{2, 1, 3},
			expected1: 10,
			expected2: 3544142,
		},
		"123": {
			in:        []int{1, 2, 3},
			expected1: 27,
			expected2: 261214,
		},
		"231": {
			in:        []int{2, 3, 1},
			expected1: 78,
			expected2: 6895259,
		},
		"321": {
			in:        []int{3, 2, 1},
			expected1: 438,
			expected2: 18,
		},
		"312": {
			in:        []int{3, 1, 2},
			expected1: 1836,
			expected2: 362,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s, err := New(tc.in, 2020, 30000000).Solve()
			assert.NoError(t, err)
			assert.Contains(t, s, fmt.Sprintf("2020 is %d and 30000000 is %d.", tc.expected1, tc.expected2))
		})
	}
}
