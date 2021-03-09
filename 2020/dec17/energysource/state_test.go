package energysource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_Activate(t *testing.T) {
	tt := map[string]struct {
		activeCells []coord3D
		min, max    coord3D
	}{
		"one": {
			activeCells: []coord3D{{1, 2, 3}},
			min:         coord3D{0, 0, 0},
			max:         coord3D{1, 2, 3},
		},
		"some": {
			activeCells: []coord3D{
				{1, 2, 3},
				{-1, 0, 2},
				{1, 2, 3},
				{-1, -2, 3},
				{1, 2, 3},
			},
			min: coord3D{-1, -2, 0},
			max: coord3D{1, 2, 3},
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
