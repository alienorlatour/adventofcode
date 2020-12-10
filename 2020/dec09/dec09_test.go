package dec09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec09_Solve(t *testing.T) {
	s, err := New("testdata/in.txt", 5).Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "127")
	// assert.Contains(t, s, "62")
}
