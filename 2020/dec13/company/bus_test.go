package company

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusCompany_FirstDeparture(t *testing.T) {
	tt := map[string]struct {
		time int
		bus  Bus
		next int
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
			assert.Equal(t, tc.next, tc.bus.firstDeparture(tc.time))
		})
	}
}
