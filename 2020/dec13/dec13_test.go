package dec13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec13_Solve(t *testing.T) {
	s, err := New("testdata/in.txt", 0).Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "Bus 59 departs in 5 minutes: 295. Next rule validation is 1068781.")
}
