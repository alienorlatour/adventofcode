package source4d

// State is an implementation of the Experimental Energy Source
// It represents the state of the Source at a given point in the iterative dimension (think time, for example)
type State struct {
	active map[coord4D]struct{}
	min    coord4D
	max    coord4D
}

type coord4D struct {
	x, y, z, w int
}

func New() *State {
	return &State{active: make(map[coord4D]struct{})}
}

// Activate sets the given cell to Active
func (s *State) Activate(x, y, z, w int) {
	c := coord4D{x: x, y: y, z: z, w:w}
	s.activate(c)
}

func (s *State) activate(c coord4D) {
	s.active[c] = struct{}{}
	s.updateBoundaries(c)
}

func (s *State) updateBoundaries(c coord4D) {
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
	if c.w < s.min.w {
		s.min.w = c.w
	}
	if c.w > s.max.w {
		s.max.w = c.w
	}
}

// CountActiveCells returns the number of active cells after a given number of iterations
func (s *State) CountActiveCellsAfter(times int) int {
	n := s
	for i := 0; i < times; i++ {
		n = n.iterate()
	}
	return len(n.active)
}
