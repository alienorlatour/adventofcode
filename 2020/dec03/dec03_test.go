package dec03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec03_Solve(t *testing.T) {
	s, err := New("testinput.txt", [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}).Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "336")
}
