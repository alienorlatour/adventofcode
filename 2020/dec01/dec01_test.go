package dec01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	// nominal
	s, err := dec01{"input.txt"}.Solve()
	assert.NoError(t, err)
	assert.Contains(t, s, "1020036")

	// no such file
	_, err = dec01{"noSuchFile"}.Solve()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file")

	// no match
	s, err = dec01{"invalidtest.txt"}.Solve()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no match found")
}
