package dec01

import (
	"fmt"
	"sort"
	"strconv"

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
	top := make(topElves, 3)

	err := fileread.ReadAndApply(d.input, func(s string) error {
		if len(s) == 0 {
			// nextElf
			top.saveIfMax(count)
			count = 0
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			count += i
		}
		return nil
	})

	top.saveIfMax(count)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", top.sum()), nil
}

type topElves []int

func (te topElves) saveIfMax(i int) {
	if te[0] < i {
		te[0] = i
		sort.Ints(te)
	}
}

func (te topElves) sum() int {
	var s int
	for _, i := range te {
		s += i
	}
	return s
}
