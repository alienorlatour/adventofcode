package source4d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_Activate(t *testing.T) {
	tt := map[string]struct {
		activeCells []coord4D
		min, max    coord4D
	}{
		"one": {
			activeCells: []coord4D{{1, 2, 3, 0}},
			min:         coord4D{0, 0, 0, 0},
			max:         coord4D{1, 2, 3, 0},
		},
		"some": {
			activeCells: []coord4D{
				{1, 2, 3, 0},
				{-1, 0, 2, 0},
				{1, 2, 3, -2},
				{-1, -2, 3, 0},
				{1, 2, 3,5},
			},
			min: coord4D{-1, -2, 0, -2},
			max: coord4D{1, 2, 3, 5},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := New()
			for _, c := range tc.activeCells {
				s.activate(c)
			}
			_, isActive := s.active[tc.activeCells[0]]
			assert.True(t, isActive)
			assert.Equal(t, tc.min, s.min)
			assert.Equal(t, tc.max, s.max)
		})
	}
}
