package dec05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBP(t *testing.T) {
	tt := map[string]struct {
		row         int
		column      int
		seatID      int
		expectedErr string
	}{
		"BFFFBBFRRR": {
			70, 7, 567, "",
		},
		"FFFBBBFRRR": {
			14, 7, 119, "",
		},
		"BBFFBBFRLL": {
			102, 4, 820, "",
		},
		"BONJOUR": {
			0, 0, 0, "invalid",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r, c, e := parseBP(name)
			if tc.expectedErr != "" {
				assert.Error(t, e)
				assert.Contains(t, e.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, e)
				assert.Equal(t, tc.row, r)
				assert.Equal(t, tc.column, c)
				assert.Equal(t, tc.seatID, BoardingPass{
					row:    r,
					column: c,
				}.SeatID()) // this should be tested in its own UT
			}
		})
	}
}
