package newmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompute(t *testing.T) {
	tt := map[string]int{
		"3":             3,
		"3 + 2":         5,
		"3 + 2 + 1 + 0": 6,
		"3 + (2 * 2)":   12,
		"(3 + 2) * 2":   10,
		"(5 + 5) * (4 + 4)": 80,
		// examples from the STORY
		"1 + 2 * 3 + 4 * 5 + 6":                           71,
		"2 * 3 + (4 * 5)":                                 26,
		"5 + (8 * 3 + 9 + 3 * 4 * 3)":                     437,
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))":       12240,
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2": 13632,
	}

	for name, result := range tt {
		t.Run(name, func(t *testing.T) {
			res, err := Compute(name)
			require.NoError(t, err)
			assert.Equal(t, result, res)
		})
	}
}
