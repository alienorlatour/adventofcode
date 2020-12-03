package dec05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec05_Solve(t *testing.T) {
	s, err := New("testinput.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "336")
}
