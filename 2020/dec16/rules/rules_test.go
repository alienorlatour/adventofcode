package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRules_AreValid(t *testing.T) {
	rules := Rules{
		rules: []Rule{
			{ranges: []ruleRange{{1, 3}, {5, 7}}},
			{ranges: []ruleRange{{6, 11}, {33, 44}}},
			{ranges: []ruleRange{{13, 40}, {45, 50}}},
		},
	}
	for _, tc := range []struct {
		values   []int
		expected int
	}{
		{[]int{7,3,47}, 0},
		{[]int{40,4,50}, 4},
		{[]int{55,2,20}, 55},
		{[]int{38,6,12}, 12},
	} {
		assert.Equal(t, rules.InvalidSum(tc.values), tc.expected)
	}
}
