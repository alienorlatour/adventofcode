package rules

// NewRule is an ugly way to create a rule
func NewRule(min1, max1, min2, max2 int, isDeparture bool) (Rule, error) {
	// todo validate min < max
	return Rule{
		ranges: []ruleRange{
			{min1, max1},
			{min2, max2},
		},
		isDeparture: isDeparture,
		lock: -1,
	}, nil
}

// Rule defines one rule
type Rule struct {
	ranges      []ruleRange
	positions   int
	maxMask     int
	isDeparture bool
	lock        int // lock is set to the position once this rule has found its place
}

// IsValid returns whether the value matches the rule
func (ru *Rule) IsValid(v int) bool {
	for _, rg := range ru.ranges {
		if rg.isValid(v) {
			return true
		}
	}
	return false
}

// SieveTicket sieves possible possitions by using the values on a given ticket
func (ru *Rule) SieveTicket(ticket Ticket) {
	for i, v := range ticket {
		if !ru.IsValid(v) {
			ru.positions &= (1 << i) ^ ru.maxMask
		}
	}
}

// TryLock tries to lock this rule in a given position. If mor ethan one position is still possible, then nothing is returned.
func (ru *Rule) TryLock() (int, bool) {
	if ru.lock >= 0 {
		return ru.lock, false
	}
	// if there is only one position left
	if ru.positions&(ru.positions-1) == 0 {
		for ru.positions > 0 {
			ru.lock ++
			ru.positions >>= 1
		}
		return ru.lock, true
	}
	return 0, false
}

type ruleRange struct {
	min, max int
}

func (rr *ruleRange) isValid(v int) bool {
	return rr.min <= v && rr.max >= v
}
