package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindMatch2(t *testing.T) {
	tt := map[string]struct {
		lines          []int
		needle         int
		expectedErr    string
		expectedValues [2]int
	}{
		"original example": {
			lines:          []int{1721, 979, 366, 299, 675, 1456},
			expectedValues: [2]int{1721, 299},
			needle:         2020,
		},
		"no match": {
			lines:       []int{1, 2, 3, 4, 5, 6},
			expectedErr: "no match found",
			needle:      202,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			l, m, err := FindMatch2(tc.lines, tc.needle)
			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err.Error())
			}
			assert.Equal(t, tc.expectedValues[0], l)
			assert.Equal(t, tc.expectedValues[1], m)
		})
	}
}
