package energysource

import (
	"fmt"
)

// State is an implementation of the Experimental Energy Source
// It represents the state of the Source at a given point in the iterative dimension (think time, for example)
type State struct {
	active map[coord3D]struct{}
	min    coord3D
	max    coord3D
}

type coord3D struct {
	x, y, z int
}

func New() *State {
	return &State{active: make(map[coord3D]struct{})}
}

// Activate sets the given cell to Active
func (s *State) Activate(x, y, z int) *State {
	c := coord3D{x: x, y: y, z: z}
	s.activate(c)
	return s
}

func (s *State) activate(c coord3D) {
	s.active[c] = struct{}{}
	s.updateBoundaries(c)
}

func (s *State) updateBoundaries(c coord3D) {
	if c.x < s.min.x {
		s.min.x = c.x
	}
	if c.x > s.max.x {
		s.max.x = c.x
	}
	if c.y < s.min.y {
		s.min.y = c.y
	}
	if c.y > s.max.y {
		s.max.y = c.y
	}
	if c.z < s.min.z {
		s.min.z = c.z
	}
	if c.z > s.max.z {
		s.max.z = c.z
	}
}

// CountActiveCells returns the number of active cells
func (s *State) CountActiveCells() int {
	return len(s.active)
}

// String returns a representation of the State's cells
func (s *State) String() string {
	str := "==== state ===="
	for z := s.min.z - 1; z <= s.max.z+1; z++ {
		str += fmt.Sprintln("\nz = ", z)
		for y := s.min.y - 1; y <= s.max.y+1; y++ {
			for x := s.min.x - 1; x <= s.max.x+1; x++ {
				c := coord3D{x, y, z}
				_, active := s.active[c]
				if active {
					str += "#"
				} else {
					str += "."
				}
			}
			str += "\n"
		}
	}
	return str
}
