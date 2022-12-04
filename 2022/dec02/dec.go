package dec02

import (
	"fmt"
	"strings"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 25
func New(input string) doors.Solver {
	return dec25{input}
}

type dec25 struct {
	input string
}

// Solve the day's problem
func (d dec25) Solve() (string, error) {
	var count int

	err := fileread.ReadAndApply(d.input, func(s string) error {
		opponent, outcome := split(s)

		me := mygame(opponent, outcome)

		count += score(opponent, me)
		return nil
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

func mygame(opponent RPS, outcome string) RPS {
	offset := int(outcome[0]) - 'Y'
	me := opponent + RPS(offset)
	me %= 3
	if me == 0 {
		me = 3
	}
	return me
}

func score(opponent, me RPS) int {
	pnts := points(opponent, me)

	return int(me) + pnts
}

func points(opponent RPS, me RPS) int {
	score := int(me + 4 - opponent)
	score %= 3
	score *= 3
	return score
}

func split(s string) (RPS, string) {
	spl := strings.Split(s, " ")
	return toRPS(spl[0]), spl[1]
}

func toRPS(s string) RPS {
	switch s {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	}
	return 0
}

type RPS int

const (
	rock RPS = iota + 1
	paper
	scissors
)
