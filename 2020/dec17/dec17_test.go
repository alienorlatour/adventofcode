package dec17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec17_Solve(t *testing.T) {
	s, err := New("testdata/in.txt", 6).Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "6 iterations, we have 112 active") // part 1
	assert.Contains(t, s, "6 iterations, we have 848 active")
}
