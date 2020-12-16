package dec10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec10_Solve(t *testing.T) {
	s, err := New("testdata/in.txt").Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, " 8 arrangements")

	s, err = New("testdata/larger-in.txt").Solve()
	assert.NoError(t, err)
	// assert.Contains(t, s, "220")
	assert.Contains(t, s, " 19208 ")
}

func TestCountDifferences(t *testing.T) {
	in := []int{
		1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19,
	}
	assert.Equal(t, differenciator{counters: map[int]int{
		1: 7, 3: 5,
	}}, countDifferences(in))
}
