package compiler

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
	"github.com/stretchr/testify/assert"
)

func TestCompiler_ParseLine(t *testing.T) {
	tt := map[string]struct {
		line     string
		expected ship.Instruction
		err      string
	}{
		"N": {
			line:     "N10",
			expected: ship.Instruction{Function: ship.North, Value: 10},
		},
		"S": {
			line:     "S1",
			expected: ship.Instruction{Function: ship.South, Value: 1},
		},
		"E": {
			line:     "E10",
			expected: ship.Instruction{Function: ship.East, Value: 10},
		},
		"W": {
			line:     "W526",
			expected: ship.Instruction{Function: ship.West, Value: 526},
		},
		"R": {
			line:     "R90",
			expected: ship.Instruction{Function: ship.Star, Value: 90},
		},
		"L": {
			line:     "L270",
			expected: ship.Instruction{Function: ship.Port, Value: 270},
		},
		"F": {
			line:     "F10",
			expected: ship.Instruction{Function: ship.Forward, Value: 10},
		},
		"invalid command": {
			line: "K3",
			err:  "invalid instruction",
		},
		"invalid value": {
			line: "FAR",
			err: "value",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			c := Compiler{}
			err := c.ParseLine(tc.line)
			if tc.err == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, c.code[0])
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.err)
			}
		})
	}
}
