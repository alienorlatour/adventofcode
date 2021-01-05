package bus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusCompany_FirstDeparture(t *testing.T) {
	tt := map[string]struct {
		time int64
		bus  Bus
		next int64
	}{
		"nominal": {
			time: 1,
			bus:  5,
			next: 5,
		},
		"now": {
			time: 6,
			bus:  6,
			next: 6,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.next, tc.bus.FirstDeparture(tc.time))
		})
	}
}
