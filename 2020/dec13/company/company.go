package company

import (
	"strconv"
	"strings"
)

type BusCompany interface {
	// returns the number of the first bus to depart after the given time
	FirstDeparture(time int) (busID Bus, waitingTime int)
}

func NewCompany(buses string) BusCompany {
	busLines := make([]Bus, 0)
	for _, b := range strings.Split(buses, ",") {
		if b != "x" {
			i, err := strconv.Atoi(b)
			if err == nil { // todo manage error if any
				busLines = append(busLines, Bus(i))
			}
		}
	}
	return &busCompany{
		buses: busLines,
	}
}

type busCompany struct {
	buses []Bus
}

func (c busCompany) FirstDeparture(time int) (Bus, int) {
	best := c.buses[0]
	next := best.firstDeparture(time)
	for _, b := range c.buses[1:] {
		n := b.firstDeparture(time)
		if n < next {
			best = b
			next = n
		}
	}
	return best, next
}
