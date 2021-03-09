package source4d

// iterate once and return a new State
func (s *State) iterate() *State {
	next := New()
	for w := s.min.w - 1; w <= s.max.w+1; w++ {
		for z := s.min.z - 1; z <= s.max.z+1; z++ {
			for y := s.min.y - 1; y <= s.max.y+1; y++ {
				for x := s.min.x - 1; x <= s.max.x+1; x++ {
					c := coord4D{x, y, z, w}
					active := s.nextCellState(c)
					if active {
						next.Activate(x, y, z, w)
					}
				}
			}
		}
	}
	return next
}

// nextCellState returns the activation (or dis- thereof) of a cell: it returns the next state of this cell
func (s *State) nextCellState(c coord4D) bool {
	_, isActive := s.active[c]
	activeNeighbours := s.countNeighbours(c)
	// If a cube is active and exactly 2 or 3 of its neighbours are also active, the cube remains active. Otherwise, the cube becomes inactive.
	if isActive && (activeNeighbours == 2 || activeNeighbours == 3) {
		return true
	}
	// If a cube is inactive but exactly 3 of its neighbours are active, the cube becomes active. Otherwise, the cube remains inactive.
	return !isActive && activeNeighbours == 3
}

// countNeighbours simply returns the number of active neighbours for a given cell
func (s *State) countNeighbours(c coord4D) int {
	var activeNeighbours int
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					if x == c.x && y == c.y && z == c.z && w == c.w {
						continue // yuck, really...
					}
					_, ok := s.active[coord4D{x, y, z, w}]
					if ok {
						activeNeighbours++
					}
				}
			}
		}
	}
	return activeNeighbours
}
