package dec10

import (
	"fmt"
	"sort"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 9
func New(input string) doors.Solver {
	return dec10{input}
}

type dec10 struct {
	input string
}

// Solve the day's problem
func (d dec10) Solve() (string, error) {
	adapters, err := fileread.ReadInts(d.input)
	if err != nil {
		return "", err
	}

	// add the initial element
	adapters = append(adapters, 0)

	sort.Ints(adapters)

	count := countArrangements(adapters)

	return fmt.Sprintf("found %d arrangements", count), nil
}

// countDifferences for part 1
// then solve using diff.counters[1]*diff.counters[3]
func countDifferences(adapters []int) differenciator {
	diff := differenciator{
		counters: make(map[int]int),
	}
	// count initial adapter starting from 0
	// todo danger here as well
	diff.counters[adapters[0]] = 1
	for i := range adapters[:len(adapters)-1] {
		dfrce := adapters[i+1] - adapters[i]
		_, ok := diff.counters[dfrce]
		if !ok {
			diff.counters[dfrce] = 1
		} else {
			diff.counters[dfrce]++
		}
	}
	// count the device's difference
	diff.counters[3]++
	return diff
}

type differenciator struct {
	counters map[int]int
}
