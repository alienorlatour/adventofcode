package rules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule_MatchTicket(t *testing.T) {
	tt := []struct {
		rule     Rule
		ticket   Ticket
		expected int
	}{
		{
			rule:     Rule{positions: 0b1111, maxMask: 0b1111, ranges: []ruleRange{{0, 3}, {6, 9}}},
			ticket:   []int{4, 2, 4, 7},
			expected: 0b1010,
		},
		{
			rule:     Rule{positions: 0b11111, maxMask: 0b11111, ranges: []ruleRange{{0, 3}, {6, 9}}},
			ticket:   []int{0, 4, 8, 1, 2},
			expected: 0b11101,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.ticket), func(t *testing.T) {
			tc.rule.SieveTicket(tc.ticket)
			assert.Equal(t, tc.expected, tc.rule.positions)
		})
	}
}

func TestRule_TryLock(t *testing.T) {
	tt := map[string]struct {
		rule     Rule
		expected int
		ok       bool
	}{
		"0": {
			rule:     Rule{positions: 1, lock: -1},
			expected: 0,
			ok:       true,
		},
		"3": {
			rule:     Rule{positions: 0b1000, lock: -1},
			expected: 3,
			ok:       true,
		},
		"already locked": {
			rule:     Rule{lock: 2},
			expected: 2,
			ok:       false,
		},
		"no lock": {
			rule:     Rule{positions: 0b101, lock: -1},
			expected: 0,
			ok:       false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			i, ok := tc.rule.TryLock()
			assert.Equal(t, tc.expected, i)
			assert.Equal(t, tc.ok, ok)
		})
	}
}
