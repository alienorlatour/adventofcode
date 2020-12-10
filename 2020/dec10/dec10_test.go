package dec10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec10_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "")
}
