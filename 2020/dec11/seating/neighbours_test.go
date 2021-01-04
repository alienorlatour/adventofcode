package seating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_min(t *testing.T) {
	assert.Equal(t, 0, min(0, 8))
	assert.Equal(t, 8, min(8, 8))
	assert.Equal(t, -5, min(-5, 8))
	assert.Equal(t, 0, min(8, 0))
	assert.Equal(t, -5, min(0, -5))
}

func Test_maxMinusOne(t *testing.T) {
	assert.Equal(t, uint(7), maxMinusOne(0, 8))
	assert.Equal(t, uint(8), maxMinusOne(8, 9))
	assert.Equal(t, uint(8), maxMinusOne(8, 0))
}
