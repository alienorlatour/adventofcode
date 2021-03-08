package dec16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ablqk/adventofcode/2020/dec16/rules"
	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

const departureRule = "departure"

// New instance of the Door for December 16
func New(rules, tickets string, mine string) doors.Solver {
	return dec16{rules: rules, tickets: tickets, mine: mine}
}

type dec16 struct {
	rules   string
	tickets string
	mine    string
}

// Solve the day's problem
func (d dec16) Solve() (string, error) {
	// parse rules
	var r rules.Rules
	err := fileread.ReadAndApply(d.rules, func(s string) error {
		rule, err := parseRule(s)
		if err != nil {
			return err
		}
		r.AddRule(rule)
		return nil
	})
	if err != nil {
		return "", err
	}
	r.Init()

	// parse tickets
	err = fileread.ReadAndApply(d.tickets, func(s string) error {
		ticket, err := parseTicket(s)
		if err != nil {
			return err
		}
		r.SieveTicket(ticket)

		return nil
	})
	if err != nil {
		return "", err
	}

	mine, err := parseTicket(d.mine)
	if err != nil {
		return "", err
	}
	departures := r.Departures(mine)

	return fmt.Sprintf("Departure values multiply up to %d.", departures), nil
}

func parseTicket(s string) (rules.Ticket, error) {
	sp := strings.Split(s, ",")
	t := make([]int, len(sp))
	var err error
	for i, v := range sp {
		t[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid line format, cannot parse ticket: %s", s)
		}
	}
	return t, nil
}

func parseRule(s string) (rules.Rule, error) {
	reg := regexp.MustCompile(".*: ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	res := reg.FindStringSubmatch(s)
	if len(res) != 5 {
		return rules.Rule{}, fmt.Errorf("invalid line format, cannot parse rule: %s", s)
	}
	min1, err := strconv.Atoi(res[1])
	if err != nil {
		return rules.Rule{}, fmt.Errorf("invalid line format, cannot parse first min: %s", s)
	}
	max1, err := strconv.Atoi(res[2])
	if err != nil {
		return rules.Rule{}, fmt.Errorf("invalid line format, cannot parse first max: %s", s)
	}
	min2, err := strconv.Atoi(res[3])
	if err != nil {
		return rules.Rule{}, fmt.Errorf("invalid line format, cannot parse second min: %s", s)
	}
	max2, err := strconv.Atoi(res[4])
	if err != nil {
		return rules.Rule{}, fmt.Errorf("invalid line format, cannot parse second max: %s", s)
	}
	return rules.NewRule(min1, max1, min2, max2, s[:len(departureRule)] == departureRule)
}
