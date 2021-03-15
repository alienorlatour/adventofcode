package newmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiplication_compute(t *testing.T) {
	tt := map[string]struct {
		a        multiplication
		expected int
	}{
		"6*3 = 18": {
			a: multiplication{
				left:  scalar{3},
				right: scalar{6},
			},
			expected: 18,
		},
		"2*3*7 = 42": {
			a: multiplication{
				left:  scalar{2},
				right: multiplication{scalar{3}, scalar{7}},
			},
			expected: 42,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.a.compute())
		})
	}
}
