package seating

import (
	"fmt"
)

type Looper struct {
	arr           Arrangement
	nextRoomState [][]square
	iterations    int
}

func NewLooper(a Arrangement) Looper {
	return Looper{
		arr: a,
	}
}

func (l *Looper) Stabilise() int {
	changed := l.iterate()
	fmt.Print("Stabilising...\n")
	for changed {
		fmt.Print(".")
		changed = l.iterate()
	}
	fmt.Println()
	return l.arr.CountPeople()
}

func (l *Looper) iterate() bool {
	var changed bool
	l.nextRoomState = l.arr.ForkRoom()
	for i := uint(0); i < l.arr.width; i++ {
		for j := uint(0); j < l.arr.height; j++ {
			c := l.sitOneSquare(i, j)
			changed = changed || c
		}
	}
	l.arr.room = l.nextRoomState
	l.iterations++
	if changed && l.iterations > 10000 {
		// stop.... something is wrong
		fmt.Println("STOP")
		return false
	}
	return changed
}

// sitOneSquare checks whether the seat in one position should be vacant or not at the next iteration
// it modifies the nextRoomState
// RULES:
// - (1) If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// - (2) If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
// - (3) Otherwise, the seat's state does not change.
func (l *Looper) sitOneSquare(x, y uint) bool {
	sq := l.nextRoomState[y][x]
	if !sq.hasSeat {
		return false
	}
	nei := l.arr.neighboursForOneSquare(x, y)
	// count the neighbouring people and sit down or up accordingly
	var people int
	for _, c := range nei {
		if l.arr.square(c).hasPerson {
			people++
		}
	}
	if !sq.hasPerson && people == 0 {
		// RULE NUMBER 1
		l.nextRoomState[y][x].hasPerson = true
		return true
	}
	if sq.hasPerson && people >= 4 {
		// RULE NUMBER 2
		l.nextRoomState[y][x].hasPerson = false
		return true
	}
	// RULE NUMBER 3
	return false
}
