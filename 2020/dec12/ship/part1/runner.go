package part1

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
)

// New ship runner
// direction is indicated like a clock: north is 0, west is 270
func New(code []ship.Instruction, directionAngle int) ship.Runner {
	return &runner{code: code, figurehead: directionAngle}
}

type runner struct {
	code       []ship.Instruction
	lat, long  int
	figurehead int
}

// Latitude returns the ship's latitude
func (r *runner) Latitude() int {
	return r.lat
}

// Longitude returns the ship's longitude
func (r *runner) Longitude() int {
	return r.long
}

// Run the code
func (r *runner) Run() error {
	for i, line := range r.code {
		err := r.execute(line)
		if err != nil {
			return fmt.Errorf("error at line %d: %s", i, err.Error())
		}
	}
	return nil
}

func (r *runner) execute(line ship.Instruction) error {
	switch line.Function {
	// Let's assume Mercator-like coordinates
	case ship.North:
		r.lat += line.Value
	case ship.South:
		r.lat -= line.Value
	case ship.East:
		r.long += line.Value
	case ship.West:
		r.long -= line.Value
	case ship.Star:
		r.figurehead = (r.figurehead + line.Value + 360) % 360
	case ship.Port:
		r.figurehead = (r.figurehead - line.Value + 360) % 360
	case ship.Forward:
		err := r.forward(line.Value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown instruction %d", line.Function)
	}
	return nil
}

func (r *runner) forward(value int) error {
	// this could be more elegant...
	switch r.figurehead {
	case 0:
		r.lat += value
	case 180:
		r.lat -= value
	case 90:
		r.long += value
	case 270:
		r.long -= value
	default:
		return fmt.Errorf("unknown direction %d", r.figurehead)
	}
	return nil
}
