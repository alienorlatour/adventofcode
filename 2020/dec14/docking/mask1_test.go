package docking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMask(t *testing.T) {
	mask, err := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0XX")
	assert.NoError(t, err)
	assert.Equal(t, Mask1{ones: 128, zeros: maxUInt36 - 4}, mask)
}
