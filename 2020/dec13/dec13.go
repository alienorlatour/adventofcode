package dec13

import (
	"fmt"
	"strconv"

	"github.com/ablqk/adventofcode/2020/dec13/company"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 13
func New(input string) doors.Solver {
	return dec13{input}
}

type dec13 struct {
	input string
}

// Solve the day's problem
func (d dec13) Solve() (string, error) {
	time, comp, err := parse(d.input)
	if err != nil {
		return "", err
	}
	bus, busTime := comp.FirstDeparture(time)
	waiting := busTime - time
	return fmt.Sprintf("Bus %d departs in %d minutes: %d.", bus, waiting, int(bus) * waiting), nil
}

func parse(input string) (int, company.BusCompany, error) {
	s, err := fileread.ReadStrings(input)
	if err != nil {
		return 0, nil, err
	}
	if len(s) != 2 {
		return 0, nil, fmt.Errorf("invalid input file: need 2 lines, got %d", len(s))
	}
	time, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, nil, err
	}
	comp := company.NewCompany(s[1])
	return time, comp, nil
}
