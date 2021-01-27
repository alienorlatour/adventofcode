package docking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDockingValue2(t *testing.T) {
	m, err := NewMask2("000000000000000000000000000000X1001X")
	require.NoError(t, err)
	assert.Equal(t, Mask2{ones: 18, exes: 33, floaterCount: 2}, m)
}

func TestMask2_Floater(t *testing.T) {
	m := Mask2{
		exes: 141,
		ones: 16,
	}
	assert.Equal(t, Value(213), m.Floater(64, 11))
	assert.Equal(t, Value(85), m.Floater(64, 3))
	assert.Equal(t, Value(64+16), m.Floater(64, 0))

	m = Mask2{
		ones: 18,
		exes: 33,
	}
	assert.Equal(t, Value(26), m.Floater(42, 0))
	assert.Equal(t, Value(27), m.Floater(42, 1))
	assert.Equal(t, Value(58), m.Floater(42, 2))
	assert.Equal(t, Value(59), m.Floater(42, 3))
}
