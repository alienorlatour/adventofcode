package part2

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
)

func New(code []ship.Instruction, waypoinyLat, waypointLong int) ship.Runner {
	return &runner{
		code:     code,
		waypoint: coordinate{long: waypointLong, lat: waypoinyLat},
	}
}

type runner struct {
	code     []ship.Instruction
	ship     coordinate // 0,0 is the ship's starting position
	waypoint coordinate // this is a vector of the ship's bearing
}

// coordinate represents a point or a vector on the map
type coordinate struct {
	lat, long int
}

// Longitude returns the ship's longitude
func (r *runner) Longitude() int {
	return r.ship.long
}

// Latitude returns the ship's latitude
func (r *runner) Latitude() int {
	return r.ship.lat
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
		r.waypoint.lat += line.Value
	case ship.South:
		r.waypoint.lat -= line.Value
	case ship.East:
		r.waypoint.long += line.Value
	case ship.West:
		r.waypoint.long -= line.Value
	case ship.Star:
		r.turn(line.Value)
	case ship.Port:
		r.turn(-line.Value)
	case ship.Forward:
		r.ship.lat += r.waypoint.lat * line.Value
		r.ship.long += r.waypoint.long * line.Value
	default:
		return fmt.Errorf("unknown instruction %d", line.Function)
	}
	return nil
}

func (r *runner) turn(angle int) {
	// Let's only use angles between 0 and 360
	angle = (angle + 360) % 360
	// Someone clever suggested using complex numbers but it's far above what my head can do it this time of the night,
	// so let's do it the old way
	// Thank you, internet https://math.stackexchange.com/questions/1330161/how-to-rotate-points-through-90-degree
	for angle > 0 {
		lat, lon := r.waypoint.lat, r.waypoint.long
		r.waypoint.lat = -lon
		r.waypoint.long = lat
		angle -= 90
	}
}
