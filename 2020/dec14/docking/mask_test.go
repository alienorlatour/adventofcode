package docking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMask(t *testing.T) {
	mask, err := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0XX")
	assert.NoError(t, err)
	assert.Equal(t, Mask{ones: 128, zeros: maxMask - 4}, mask)
}
