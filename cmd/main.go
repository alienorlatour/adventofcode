package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ablqk/adventofcode/2020/dec01"
	"github.com/ablqk/adventofcode/2020/dec02"
	"github.com/ablqk/adventofcode/2020/dec03"
	"github.com/ablqk/adventofcode/2020/dec04"
	"github.com/ablqk/adventofcode/2020/dec05"
	"github.com/ablqk/adventofcode/2020/dec06"
	"github.com/ablqk/adventofcode/doors"
)

const lastDay = 6

func main() {
	// -d for the day of December we want to play
	d := flag.Int("day", lastDay, "the number of the day you want to play")
	flag.Parse()
	if d == nil {
		d = new(int)
		*d = lastDay
	}

	var o doors.Solver
	switch *d {
	case 1:
		o = dec01.New("2020/dec01/input.txt")
	case 2:
		o = dec02.New("2020/dec02/input.txt")
	case 3:
		o = dec03.New("2020/dec03/input.txt", [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
	case 4:
		o = dec04.New("2020/dec04/input.txt")
	case 5:
		o = dec05.New("2020/dec05/input.txt")
	case 6:
		o = dec06.New("2020/dec06/input.txt")
	// ===== ready for tomorrow =====
	default:
		log.Fatalf("unsupported day of the month %d", *d)
		return
	}
	s, err := o.Solve()
	if err != nil {
		log.Fatalf("could not compute the result for day %d: %s", *d, err.Error())
	}
	fmt.Println("====== \n  RESULT:\n  ", s, "\n======")
}
