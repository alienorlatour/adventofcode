package source4d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_iterate(t *testing.T) {
	tt := map[string]struct {
		activeCells         []coord4D
		times               int
		expectedActiveCells int
	}{
		// example from the STORY, before iterating
		"stagnant": {
			activeCells: []coord4D{
				{0, -1, 0, 0},
				{1, 0, 0, 0},
				{-1, 1, 0, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			times:               0,
			expectedActiveCells: 5,
		},
		// example from the STORY, one iteration
		"one iteration": {
			activeCells: []coord4D{
				{0, -1, 0, 0},
				{1, 0, 0, 0},
				{-1, 1, 0, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			times:               1,
			expectedActiveCells: 29,
		},
		// example from the STORY, full
		"2 iterations": {
			activeCells: []coord4D{
				{0, -1, 0, 0},
				{1, 0, 0, 0},
				{-1, 1, 0, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			times:               2,
			expectedActiveCells: 60,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := New()
			for _, c := range tc.activeCells {
				s.activate(c)
			}
			assert.Equal(t, tc.expectedActiveCells, s.CountActiveCellsAfter(tc.times))
		})
	}
}
