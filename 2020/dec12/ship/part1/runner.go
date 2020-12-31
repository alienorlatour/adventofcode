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

func (r *runner) Lattitude() int {
	return r.lat
}

func (r *runner) Longitude() int {
	return r.long
}

func (r *runner) Run() error {
	for i, line := range r.code {
		switch line.Function {
		// Let's assume Mercator-like coordinates
		case ship.GoNorth:
			r.lat += line.Value
		case ship.GoSouth:
			r.lat -= line.Value
		case ship.GoEast:
			r.long += line.Value
		case ship.GoWest:
			r.long -= line.Value
		case ship.TurnRight:
			r.figurehead = (r.figurehead + line.Value + 360) % 360
		case ship.TurnLeft:
			r.figurehead = (r.figurehead - line.Value + 360) % 360
		case ship.GoForward:
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
				return fmt.Errorf("unknown direction %d at line %d", r.figurehead, i)
			}
		default:
			return fmt.Errorf("unknown instruction %d at line %d", line.Function, i)
		}
	}
	return nil
}
