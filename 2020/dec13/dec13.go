package dec13

import (
	"fmt"
	"strconv"

	"github.com/ablqk/adventofcode/2020/dec13/company"
	"github.com/ablqk/adventofcode/2020/dec13/rule"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

// New instance of the Door for December 13
func New(input string, startingAt int64) doors.Solver {
	return dec13{input, startingAt}
}

type dec13 struct {
	input   string
	startAt int64
}

// Solve the day's problem
func (d dec13) Solve() (string, error) {
	time, comp, err := parse(d.input)
	if err != nil {
		return "", err
	}
	bus, busTime := comp.FirstDeparture(time)
	waiting := busTime - time

	busRule, err := parse2(d.input)
	if err != nil {
		return "", err
	}
	validation, err := busRule.EarliestRuleValidation(d.startAt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Bus %d departs in %d minutes: %d. Next rule validation is %d.", bus, waiting, int64(bus)*waiting, validation), nil
}

func parseFile(input string) ([]string, error) {
	s, err := fileread.ReadStrings(input)
	if err != nil {
		return nil, err
	}
	if len(s) != 2 {
		return nil, fmt.Errorf("invalid input file: need 2 lines, got %d", len(s))
	}
	return s, nil
}

func parse(input string) (int64, company.BusCompany, error) {
	s, err := parseFile(input)
	if err != nil {
		return 0, nil, err
	}
	time, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, nil, err
	}
	comp := company.NewCompany(s[1])
	// todo parse as int64
	return int64(time), comp, nil
}

func parse2(input string) (rule.BusRule, error) {
	s, err := parseFile(input)
	if err != nil {
		return nil, err
	}
	comp := rule.NewRule(s[1])
	return comp, nil
}
