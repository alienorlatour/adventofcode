package dec22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec22_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "")
}
