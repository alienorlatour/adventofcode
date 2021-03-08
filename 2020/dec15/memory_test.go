package dec15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryGame_ValueAt(t *testing.T) {
	tt := map[string]struct {
		input    []int
		end      int
		expected int
	}{
		"nominal 4": {
			input:    []int{0, 3, 6},
			end:      4,
			expected: 0,
		},
		"nominal 10": {
			input:    []int{0, 3, 6},
			end:      10,
			expected: 0,
		},
		"nominal 9": {
			input:    []int{0, 3, 6},
			end:      9,
			expected: 4,
		},
		"nominal 5": {
			input:    []int{0, 3, 6},
			end:      5,
			expected: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			m := NewMemoryGame(tc.input)
			v := m.ValueAt(tc.end)
			assert.Equal(t, tc.expected, v)
		})
	}
}
