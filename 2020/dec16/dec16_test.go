package dec16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec16_Solve(t *testing.T) {
	s, err := New("testdata/rules.txt", "testdata/tickets.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "is 71.")
}
