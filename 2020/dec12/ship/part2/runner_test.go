package part2

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
	"github.com/stretchr/testify/assert"
)

func TestRunner_execute(t *testing.T) {
	tt := map[string]struct {
		line             ship.Instruction
		startShip        coordinate
		expectedShip     coordinate
		startWaypoint    coordinate
		expectedWaypoint coordinate
	}{
		"F10 north": {
			line:             ship.Instruction{Function: ship.Forward, Value: 10},
			startWaypoint:    coordinate{1, 3},
			expectedWaypoint: coordinate{1, 3},
			expectedShip:     coordinate{10, 30},
		},
		"S3": {
			line:             ship.Instruction{Function: ship.South, Value: 3},
			expectedWaypoint: coordinate{lat: -3, long: 0},
		},
		"N4": {
			line:             ship.Instruction{Function: ship.North, Value: 4},
			expectedWaypoint: coordinate{lat: 4, long: 0},
		},
		"E4": {
			line:             ship.Instruction{Function: ship.East, Value: 4},
			expectedWaypoint: coordinate{lat: 0, long: 4},
		},
		"W4": {
			line:             ship.Instruction{Function: ship.West, Value: 4},
			expectedWaypoint: coordinate{lat: 0, long: -4},
		},
		"R90": {
			line:             ship.Instruction{Function: ship.Star, Value: 90},
			startWaypoint:    coordinate{1, 3},
			expectedWaypoint: coordinate{-3, 1},
		},
		"L180": {
			line:             ship.Instruction{Function: ship.Port, Value: 180},
			startWaypoint:    coordinate{1, 3},
			expectedWaypoint: coordinate{-1, -3},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := runner{
				ship:     tc.startShip,
				waypoint: tc.startWaypoint,
			}
			err := r.execute(tc.line)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedShip, r.ship)
			assert.Equal(t, tc.expectedWaypoint, r.waypoint)
		})
	}
}
