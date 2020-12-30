package seating

// func (a *Arrangement) neighboursForOneSquarePart1(x, y uint) (neighbours []coordinate) {
// 	for j := maxMinusOne(0, y); j < min(a.height, y+2); j++ {
// 		for i := maxMinusOne(0, x); i < min(a.width, x+2); i++ {
// 			if i == x && j == y {
// 				// I am not my own neighbour
// 				continue
// 			}
// 			c := coordinate{i, j}
// 			if a.square(c).hasSeat {
// 				neighbours = append(neighbours, c)
// 			}
// 		}
// 	}
// 	return neighbours
// }

func (a *Arrangement) neighboursForOneSquarePart2(x, y int) (neighbours []coordinate) {
	for yMod := - 1; yMod <= 1; yMod++ {
		for xMod := - 1; xMod <= 1; xMod++ {
			if xMod == 0 && yMod == 0 {
				// I am not my own neighbour
				continue
			}
			n := a.neighbourInOneDir(x, y, xMod, yMod)
			if n != nil {
				neighbours = append(neighbours, *n)
			}
		}
	}
	return neighbours
}

// xMod = 1 means we go right, yMod = -1 means we go up, and so on
func (a *Arrangement) neighbourInOneDir(x, y int, xMod, yMod int) (neighbours *coordinate) {
	distance := 1
	c := coordinate{x + distance*xMod, y + distance*yMod}
	// while we can't see a seat and don't reach any of the four walls the walls
	for {
		if c.x < 0 || c.y < 0 || c.x >= a.width || c.y >= a.height {
			// that's a wall
			return nil
		}
		if a.square(c).hasSeat {
			// hello neighbour
			return &c
		}
		distance++
		c = coordinate{x + distance*xMod, y + distance*yMod}
	}
}

func min(i, j int) int {
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
