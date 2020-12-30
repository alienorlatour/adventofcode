package seating

import (
	"fmt"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

// Arrangement describes the arrangement of seats of floor tiles in the room
type Arrangement struct {
	width, height int
	room          [][]square
}

// one square in the room
type square struct {
	hasSeat   bool
	hasPerson bool
}

func (s square) String() string {
	if !s.hasSeat {
		return string(floor)
	}
	if s.hasPerson {
		return string(occupied)
	}
	return string(empty)
}

// a seat's coordinates
type coordinate struct {
	x, y int
}

func (a *Arrangement) ParseLine(s string) error {
	// validate width
	if a.width == 0 {
		a.width = int(len(s))
	} else if a.width != int(len(s)) {
		return fmt.Errorf("invalid line length, expected %d got %d", a.width, len(s))
	}
	// parse each char
	line := make([]square, a.width)
	for i, c := range []byte(s) {
		switch c {
		case floor: // meh
		case empty:
			line[i] = square{
				hasSeat: true,
			}
		case occupied:
			return fmt.Errorf("we should be starting with an empty room")
		}
	}
	a.room = append(a.room, line)
	a.height++
	return nil
}

func (a *Arrangement) square(c coordinate) *square {
	if c.x >= a.width || c.y >= a.height {
		return nil
	}
	return &a.room[c.y][c.x]
}

func (a *Arrangement) CountPeople() int {
	var count int
	for _, line := range a.room {
		for _, place := range line {
			if place.hasPerson {
				count++
			}
		}
	}
	return count
}

func (a *Arrangement) ForkRoom() [][]square {
	// gracefully copied from StackOverflow
	duplicate := make([][]square, len(a.room))
	for i := range a.room {
		duplicate[i] = make([]square, len(a.room[i]))
		copy(duplicate[i], a.room[i])
	}
	return duplicate
}

func (a *Arrangement) String() string {
	var s string
	for i := 0; i < a.height; i++ {
		for _, place := range a.room[i] {
			s += place.String()
		}
		s += "\n"
	}
	return s
}
