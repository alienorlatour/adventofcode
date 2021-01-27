package dec14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec14_Solve(t *testing.T) {
	s, err := New("testdata/in2.txt").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "addition is 165.") // part 1
	assert.Contains(t, s, "addition is 208.")
}
