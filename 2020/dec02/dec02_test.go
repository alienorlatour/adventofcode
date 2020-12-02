package dec02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec02_Solve(t *testing.T) {
	s, err := New("testdata/example.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "1")
}
