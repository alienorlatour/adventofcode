package rule

import (
	"strconv"
	"strings"

	"github.com/ablqk/adventofcode/2020/dec13/bus"
)

type BusRule interface {
	// gives the earliest timestamp that validates the rule, after the given time
	EarliestRuleValidation(startTime int64) (int64, error)
}

func NewRule(buses string) BusRule {
	spl := strings.Split(buses, ",")
	busLines := make(map[int]bus.Bus)
	for i, b := range spl {
		if b != "x" {
			id, err := strconv.Atoi(b)
			if err == nil { // todo manage error if any
				busLines[i] = bus.New(id)
			}
		}
	}
	return &busRule{
		rule: busLines,
	}
}

type busRule struct {
	// index of the bus, id of the bus
	rule map[int]bus.Bus
}

func (r busRule) EarliestRuleValidation(startTime int64) (int64, error) {
	diff, biggestBus := r.biggestBus()
	// try each time the biggest bus id departs
	for time := biggestBus.FirstDeparture(startTime) - int64(diff); ; time += int64(biggestBus) {
		if r.isValid(time) {
			return time, nil
		}
	}
}

// isValid tells whether the given time validates the rule
func (r busRule) isValid(time int64) bool {
	// for each bus, check that its supposed starting time is a valid starting time
	for i, b := range r.rule {
		shouldLeaveAt := time + int64(i)
		if shouldLeaveAt%int64(b) != 0 {
			// bus r does not leave at this time
			return false
		}
	}
	return true
}

func (r busRule) biggestBus() (int, bus.Bus) {
	var biggestIndex int
	for i, b := range r.rule {
		if b > r.rule[biggestIndex] {
			biggestIndex = i
		}
	}
	return biggestIndex, r.rule[biggestIndex]
}
