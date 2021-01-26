package docking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDockingValue(t *testing.T) {
	v, err := NewDockingValue("11")
	assert.NoError(t, err)
	assert.Equal(t, value(11), v)
}

func TestValue_ApplyMask(t *testing.T) {
	m, _ := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	assert.Equal(t, value(73), value(11).ApplyMask(m))
	assert.Equal(t, value(101), value(101).ApplyMask(m))
	assert.Equal(t, value(64), value(0).ApplyMask(m))
}
