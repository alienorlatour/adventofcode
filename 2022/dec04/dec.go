package dec04

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 25
func New(input string) doors.Solver {
	return door{input}
}

type door struct {
	input string
}

// Solve the day's problem
func (d door) Solve() (string, error) {
	var count int

	err := fileread.ReadAndApply(d.input, func(s string) error {
		p1, p2 := parse(s)
		// if p1.contains(p2) || p2.contains(p1) {
		if p1.overlaps(p2) {
			count++
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

func parse(s string) (pair, pair) {
	reg := regexp.MustCompile("(\\d+)-(\\d+),(\\d+)-(\\d+)")

	separ := reg.FindStringSubmatch(s)
	numbers := make([]int, 4)

	for i, n := range separ[1:] {
		numbers[i], _ = strconv.Atoi(n)
	}

	return pair{numbers[0], numbers[1]}, pair{numbers[2], numbers[3]}
}

type pair struct {
	min, max int
}

func (p pair) contains(right pair) bool {
	return p.min <= right.min && p.max >= right.max
}

func (p pair) overlaps(right pair) bool {
	return !(p.max < right.min || p.min > right.max)
}
