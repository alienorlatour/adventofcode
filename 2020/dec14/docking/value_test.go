package docking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDockingValue(t *testing.T) {
	v, err := NewDockingValue("11")
	assert.NoError(t, err)
	assert.Equal(t, Value(11), v)
}

func TestMask1_ApplyTo(t *testing.T) {
	m, _ := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	assert.Equal(t, Value(73), m.ApplyTo(Value(11)))
	assert.Equal(t, Value(101), m.ApplyTo(Value(101)))
	assert.Equal(t, Value(64), m.ApplyTo(Value(0)))
}
