package company

import (
	"strconv"
	"strings"

	"github.com/ablqk/adventofcode/2020/dec13/bus"
)

type BusCompany interface {
	// returns the ID of the first bus to depart after the given time
	FirstDeparture(time int64) (busID bus.Bus, waitingTime int64)
}

func NewCompany(buses string) BusCompany {
	busLines := make([]bus.Bus, 0)
	for _, b := range strings.Split(buses, ",") {
		if b != "x" {
			i, err := strconv.Atoi(b)
			if err == nil { // todo manage error if any
				busLines = append(busLines, bus.New(i))
			}
		}
	}
	return &busCompany{
		buses: busLines,
	}
}

type busCompany struct {
	buses []bus.Bus
	rule  []int
}

func (c busCompany) FirstDeparture(time int64) (bus.Bus, int64) {
	best := c.buses[0]
	next := best.FirstDeparture(time)
	for _, b := range c.buses[1:] {
		n := b.FirstDeparture(time)
		if n < next {
			best = b
			next = n
		}
	}
	return best, next
}
