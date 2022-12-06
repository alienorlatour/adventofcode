package dec03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "70")
	// assert.Contains(t, s, "157")
}
