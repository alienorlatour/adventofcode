package dec12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec12_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "Manhattan distance is 25.")
	assert.Contains(t, s, "Manhattan distance is 286.")

	s, err = New("testdata/in2.txt").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "Manhattan distance is 18.")
	assert.Contains(t, s, "Manhattan distance is 462.")
}
