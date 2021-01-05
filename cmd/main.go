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
	"github.com/ablqk/adventofcode/2020/dec07"
	"github.com/ablqk/adventofcode/2020/dec08"
	"github.com/ablqk/adventofcode/2020/dec09"
	"github.com/ablqk/adventofcode/2020/dec10"
	"github.com/ablqk/adventofcode/2020/dec11"
	"github.com/ablqk/adventofcode/2020/dec12"
	"github.com/ablqk/adventofcode/2020/dec13"
	"github.com/ablqk/adventofcode/2020/dec14"
	"github.com/ablqk/adventofcode/2020/dec15"
	"github.com/ablqk/adventofcode/2020/dec16"
	"github.com/ablqk/adventofcode/doors"
)

const (
	lastDay = 16
)

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
	case 7:
		o = dec07.New("2020/dec07/input.txt", "shiny gold")
	case 8:
		o = dec08.New("2020/dec08/input.txt")
	case 9:
		o = dec09.New("2020/dec09/input.txt", 25)
	case 10:
		o = dec10.New("2020/dec10/input.txt")
	case 11:
		o = dec11.New("2020/dec11/input.txt")
	case 12:
		o = dec12.New("2020/dec12/input.txt")
	case 13:
		o = dec13.New("2020/dec13/input.txt", 471759240000000)
	case 14:
		o = dec14.New("2020/dec14/input.txt")
	case 15:
		o = dec15.New("2020/dec15/input.txt")
	case 16:
		o = dec16.New("2020/dec16/input.txt")
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
