package dec05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec05_Solve(t *testing.T) {
	// s, err := New("testdata/in.txt").Solve()
	// assert.NoError(t, err)
	// assert.Contains(t, s, "820")

	s, err := New("testdata/in2.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "5")
}
