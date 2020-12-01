package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ablqk/adventofcode/2020/dec01"
	"github.com/ablqk/adventofcode/2020/dec02"
	"github.com/ablqk/adventofcode/output"
)

const lastDay = 2

func main() {
	// -d for the day of December we want to play
	d := flag.Int("day", lastDay, "the number of the day you want to play")
	flag.Parse()
	if d == nil {
		d = new(int)
		*d = lastDay
	}

	var o output.Outputter
	switch *d {
	case 1:
		o = dec01.New("2020/dec01/input.txt")
	case 2:
		o = dec02.New()
	// ===== ready for tomorrow =====
	default:
		log.Fatalf("unsupported day of the month %d", *d)
		return
	}
	s, err := o.Output()
	if err != nil {
		log.Fatalf("could not compute the result for day %d: %s", *d, err.Error())
	}
	fmt.Println("====== \n  RESULT:\n  ", s, "\n======")
}
