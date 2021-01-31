package rules

type Rules struct {
	rules []Rule
}

func (r *Rules) AddRule(rule Rule) {
	r.rules = append(r.rules, rule)
}

func (r Rules) InvalidSum(values []int) int {
	var sum int
	for _, v := range values {
		if !r.IsValid(v) {
			sum += v
		}
	}
	return sum
}

func (r Rules) IsValid(v int) bool {
	for _, ru := range r.rules {
		if ru.IsValid(v) {
			return true
		}
	}
	return false
}

func NewRule(min1, max1, min2, max2 int) (Rule, error) {
	// todo validate min < max
	return Rule{
		ranges: []ruleRange{
			{min1, max1},
			{min2, max2},
		},
	}, nil
}

type Rule struct {
	ranges []ruleRange
}

func (ru Rule) IsValid(v int) bool {
	for _, rg := range ru.ranges {
		if rg.IsValid(v) {
			return true
		}
	}
	return false
}

type ruleRange struct {
	min, max int
}

func (rr ruleRange) IsValid(v int) bool {
	return rr.min <= v && rr.max >= v
}
