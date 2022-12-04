package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ablqk/adventofcode/2022/dec01"
	"github.com/ablqk/adventofcode/doors"
)

const (
	lastDay = 25
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
		o = dec01.New("2022/dec01/input.txt")
	// case 2:
	// 	o = dec02.New("2022/dec02/input.txt")
	// case 3:
	// 	o = dec03.New("2022/dec03/input.txt", [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
	// case 4:
	// 	o = dec04.New("2022/dec04/input.txt")
	// case 5:
	// 	o = dec05.New("2022/dec05/input.txt")
	// case 6:
	// 	o = dec06.New("2022/dec06/input.txt")
	// case 7:
	// 	o = dec07.New("2022/dec07/input.txt", "shiny gold")
	// case 8:
	// 	o = dec08.New("2022/dec08/input.txt")
	// case 9:
	// 	o = dec09.New("2022/dec09/input.txt", 25)
	// case 10:
	// 	o = dec10.New("2022/dec10/input.txt")
	// case 11:
	// 	o = dec11.New("2022/dec11/input.txt")
	// case 12:
	// 	o = dec12.New("2022/dec12/input.txt")
	// case 13:
	// 	o = dec13.New("2022/dec13/input.txt", 471759240000000)
	// case 14:
	// 	o = dec14.New("2022/dec14/input.txt")
	// case 15:
	// 	o = dec15.New([]int{9,3,1,0,8,4}, 2022, 30000000)
	// case 16:
	// 	o = dec16.New("2022/dec16/rules.txt", "2022/dec16/nearby.txt", "79,149,97,163,59,151,101,89,173,139,167,61,73,71,137,53,83,157,131,67")
	// case 17:
	// 	o = dec17.New("2022/dec17/input.txt")
	// case 18:
	// 	o = dec18.New("2022/dec18/input.txt")
	// case 19:
	// 	o = dec19.New("2022/dec19/input.txt")
	// case 20:
	// 	o = dec20.New("2022/dec20/input.txt")
	// case 21:
	// 	o = dec21.New("2022/dec21/input.txt")
	// case 22:
	// 	o = dec22.New("2022/dec22/input.txt")
	// case 23:
	// 	o = dec23.New("2022/dec23/input.txt")
	// case 24:
	// 	o = dec24.New("2022/dec24/input.txt")
	// case 25:
	// 	o = dec25.New("2022/dec25/input.txt")
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
