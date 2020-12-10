package dec07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec07_Solve(t *testing.T) {
	s, err := New("testdata/in.txt", "shiny gold").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "4")
	assert.Contains(t, s, "32")
}
