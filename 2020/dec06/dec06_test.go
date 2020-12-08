package dec06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec06_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "11") // part 1
	assert.Contains(t, s, "6")
}
