package rules

import (
	"sort"
)

// Rules can tell what rule corresponds to what field
// we use the very handy zero-value initialisation:
// posMatrix tells whether a given rule is blocked from a given position
type Rules struct {
	rules []Rule
}

// AddRule to the rules set
func (r *Rules) AddRule(rule Rule) {
	r.rules = append(r.rules, rule)
}

// Init the rules - I really dislike init methods, though
func (r *Rules) Init() {
	// as many ones as there are rules
	possibilities := 1<<len(r.rules) - 1
	for i := range r.rules {
		r.rules[i].positions = possibilities
		r.rules[i].maxMask = possibilities
	}
}

// InvalidSum sums the invalid values in the ticket
func (r Rules) InvalidSum(ticket Ticket) int {
	var sum int
	for _, v := range ticket {
		if !r.isValid(v) {
			sum += v
		}
	}
	return sum
}

func (r Rules) isValid(v int) bool {
	for _, ru := range r.rules {
		if ru.IsValid(v) {
			return true
		}
	}
	return false
}

// SieveTicket crosses out positions for the rules
func (r *Rules) SieveTicket(ticket Ticket) {
	if r.InvalidSum(ticket) != 0 {
		// ignore invalid tickets
		// we know that 0 is not among the invalid values - that's a fix for the future
		return
	}
	for i := range r.rules {
		r.rules[i].SieveTicket(ticket)
	}
}

// finaliseDepartures returns the positions of the Departure rules
func (r *Rules) finalise() []int {
	ordered := make(map[int]Rule, len(r.rules))
	// let's trust that there is a solution
	for len(ordered) < len(r.rules) {
		for i, ru := range r.rules {
			// try to lock this line (rule)
			index, ok := r.rules[i].TryLock()
			if ok {
				ordered[index] = ru
				r.lockColumn(index)
			}
			// try to lock this column (position)
			rule, ok := r.tryLockColumn(i)
			if ok {
				ordered[i] = *rule
			}
		}
	}
	var departures []int
	for i, ru := range ordered {
		if ru.isDeparture {
			departures = append(departures, i)
		}
	}
	sort.Ints(departures)
	return departures
}

// Departures basically returns the door's solution
func (r *Rules) Departures(mine Ticket) int64 {
	ordered := r.finalise()
	count := int64(1)
	for _, i := range ordered {
		count *= int64(mine[i])
	}
	return count
}

// lockColumn removes this position for all possibilities
func (r *Rules) lockColumn(index int) {
	mask := (1 << index) ^ ((1 << len(r.rules)) - 1)
	for i := range r.rules {
		r.rules[i].positions &= mask
	}
}

func (r *Rules) tryLockColumn(i int) (*Rule, bool) {
	mask := 1 << i
	var rule *Rule
	for irr, rr := range r.rules {
		// is this rule valid in the column
		if rr.positions&mask == 1 {
			if rule != nil {
				// too many possibilities
				return nil, false
			}
			rule = &r.rules[irr]
		}
	}
	if rule != nil {
		rule.lock = i
		rule.positions = mask
		r.lockColumn(i)
		return rule, true
	}
	return nil, false
}
