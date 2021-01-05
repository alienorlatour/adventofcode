package rule

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec13/bus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBusRule_EarliestRuleValidation(t *testing.T) {
	tt := map[string]struct {
		rules     map[int]bus.Bus
		validTime int64
	}{
		"7,13,x,x,59,x,31,19": {
			rules: map[int]bus.Bus{
				0: 7, 1: 13, 4: 59, 6: 31, 7: 19,
			},
			validTime: 1068781,
		},
		"17,x,13,19": {
			rules: map[int]bus.Bus{
				0: 17, 2: 13, 3: 19,
			},
			validTime: 3417,
		},
		"67,7,59,61": {
			rules: map[int]bus.Bus{
				0: 67, 1: 7, 2: 59, 3: 61,
			},
			validTime: 754018,
		},
		"67,x,7,59,61 first occurs at timestamp 779210": {
			rules: map[int]bus.Bus{
				0: 67, 2: 7, 3: 59, 4: 61,
			},
			validTime: 779210,
		},
		"67,7,x,59,61 first occurs at timestamp 1261476": {
			rules: map[int]bus.Bus{
				0: 67, 1: 7, 3: 59, 4: 61,
			},
			validTime: 1261476,
		},
		"1789,37,47,1889 first occurs at timestamp 1202161486": {
			rules: map[int]bus.Bus{
				0: 1789, 1: 37, 2: 47, 3: 1889,
			},
			validTime: 1202161486,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			time, e := busRule{tc.rules}.EarliestRuleValidation(0)
			require.NoError(t, e)
			assert.Equal(t, tc.validTime, time)
		})
	}
}
