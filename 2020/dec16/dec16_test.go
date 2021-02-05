package dec16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec16_Solve(t *testing.T) {
	s, err := New("testdata/rules.txt", "testdata/tickets.txt", "11,12,13").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "up to 132.")
}
