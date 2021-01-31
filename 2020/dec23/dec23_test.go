package dec23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec23_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "")
}
