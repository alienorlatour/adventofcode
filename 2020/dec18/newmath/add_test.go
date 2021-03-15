package newmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addition_compute(t *testing.T) {
	tt := map[string]struct {
		a        addition
		expected int
	}{
		"6+3 = 9": {
			a: addition{
				left:  scalar{3},
				right: scalar{6},
			},
			expected: 9,
		},
		"6+3+7 = 16": {
			a: addition{
				left:  scalar{6},
				right: addition{scalar{3}, scalar{7}},
			},
			expected: 16,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.a.compute())
		})
	}
}
