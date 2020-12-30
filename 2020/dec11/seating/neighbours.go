package seating

func (a *Arrangement) neighboursForOneSquare(x, y uint) (neighbours []coordinate) {
	for j := maxMinusOne(0, y); j < min(a.height, y+2); j++ {
		for i := maxMinusOne(0, x); i < min(a.width, x+2); i++ {
			if i == x && j == y {
				// I am not my own neighbour
				continue
			}
			c := coordinate{i, j}
			if a.square(c).hasSeat {
				neighbours = append(neighbours, c)
			}
		}
	}
	return neighbours
}

func min(i, j uint) uint {
	if i <= j {
		return i
	}
	return j
}

// maxMinusOne returns the maximum between x and minusOne-1, taking charge of the overflow problem
func maxMinusOne(i, minusOne uint) uint {
	if minusOne == 0 {
		// uint(0-1) overflows
		return 0
	}
	j := minusOne - 1
	if i >= j {
		return i
	}
	return j
}
