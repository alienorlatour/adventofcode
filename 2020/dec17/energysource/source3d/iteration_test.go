package source3d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState_Iterate(t *testing.T) {
	tt := map[string]struct {
		activeCells         []coord3D
		times               int
		expectedActiveCells int
	}{
		// example from the STORY, before iterating
		"stagnant": {
			activeCells: []coord3D{
				{0, -1, 0},
				{1, 0, 0},
				{-1, 1, 0},
				{0, 1, 0},
				{1, 1, 0},
			},
			times:               0,
			expectedActiveCells: 5,
		},
		// example from the STORY, one iteration
		"one iteration": {
			activeCells: []coord3D{
				{0, -1, 0},
				{1, 0, 0},
				{-1, 1, 0},
				{0, 1, 0},
				{1, 1, 0},
			},
			times:               1,
			expectedActiveCells: 11,
		},
		// example from the STORY, full
		"3 iterations": {
			activeCells: []coord3D{
				{0, -1, 0},
				{1, 0, 0},
				{-1, 1, 0},
				{0, 1, 0},
				{1, 1, 0},
			},
			times:               3,
			expectedActiveCells: 38,
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

func Test_countNeighbours(t *testing.T) {
	tt := map[string]struct {
		state    State
		coord    coord3D
		expected int
	}{
		"only you can make the darkness bright": {
			state: State{
				active: map[coord3D]struct{}{
					{0, 0, 0}: {},
				},
			},
			coord:    coord3D{0, 0, 0},
			expected: 0,
		},
		"corners": {
			state: State{
				active: map[coord3D]struct{}{
					{-1, 1, 1}:  {},
					{-1, 1, -1}: {},
					{1, 1, 1}:   {},
					{-1, -1, 1}: {},
				},
			},
			coord:    coord3D{0, 0, 0},
			expected: 4,
		},
		"story's example": {
			state: State{
				active: map[coord3D]struct{}{
					{0, -1, 0}: {},
					{1, 0, 0}:  {},
					{-1, 1, 0}: {},
					{0, 1, 0}:  {},
					{1, 1, 0}:  {},
				},
				min: coord3D{-1,-1,-1},
				max: coord3D{1,1,1},
			},
			coord:    coord3D{0, -1, 0},
			expected: 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.state.countNeighbours(tc.coord))
		})
	}
}
