package part1

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
	"github.com/stretchr/testify/assert"
)

func TestRunner_execute(t *testing.T) {
	tt := map[string]struct {
		line                      ship.Instruction
		startLat, startLong       int
		startFigurehead           int
		expectedLat, expectedLong int
		expectedFigurehead        int
	}{
		"F10 north": {
			line:            ship.Instruction{Function: ship.Forward, Value: 10},
			startFigurehead: 0, // north
			startLat:        0,
			expectedLat:     10,
		},
		"F10 south": {
			line:               ship.Instruction{Function: ship.Forward, Value: 10},
			startFigurehead:    180, // south
			expectedFigurehead: 180,
			startLat:           0,
			expectedLat:        -10,
		},
		"F10 east": {
			line:               ship.Instruction{Function: ship.Forward, Value: 10},
			startFigurehead:    90, // east
			expectedFigurehead: 90,
			expectedLong:       10,
		},
		"S3": {
			line:               ship.Instruction{Function: ship.South, Value: 3},
			startFigurehead:    270,
			expectedFigurehead: 270,
			expectedLat:        -3,
		},
		"N4": {
			line:        ship.Instruction{Function: ship.North, Value: 4},
			startLat:    3,
			expectedLat: 7,
		},
		"E4": {
			line:         ship.Instruction{Function: ship.East, Value: 4},
			startLong:    2,
			expectedLong: 6,
			startLat:     -3,
			expectedLat:  -3,
		},
		"W4": {
			line:         ship.Instruction{Function: ship.West, Value: 4},
			expectedLong: -4,
		},
		"R90": {
			line:               ship.Instruction{Function: ship.Star, Value: 90},
			startFigurehead:    270, // west
			expectedFigurehead: 0,   // north
		},
		"L180": {
			line:               ship.Instruction{Function: ship.Port, Value: 180},
			startFigurehead:    90,
			expectedFigurehead: 270,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := runner{
				lat:        tc.startLat,
				long:       tc.startLong,
				figurehead: tc.startFigurehead,
			}
			err := r.execute(tc.line)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedLat, r.Latitude())
			assert.Equal(t, tc.expectedLong, r.Longitude())
			assert.Equal(t, tc.expectedFigurehead, r.figurehead)
		})
	}
}
