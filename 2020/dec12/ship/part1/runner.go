package part1

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
)

func New(code []ship.Instruction, directionAngle int) ship.Runner {
	return &runner{code: code, figurehead: directionAngle}
}

type runner struct {
	code       []ship.Instruction
	lat, long  int
	figurehead int // north is 0, west is 270
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
		// this could be more elegant...
		switch r.figurehead {
		case 0:
			r.lat += line.Value
		case 180:
			r.lat -= line.Value
		case 90:
			r.long += line.Value
		case 270:
			r.long -= line.Value
		default:
			return fmt.Errorf("unknown direction %d", r.figurehead)
		}
	default:
		return fmt.Errorf("unknown instruction %d", line.Function)
	}
	return nil
}
