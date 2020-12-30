package dec11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec11_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "there are 37 people")
}
