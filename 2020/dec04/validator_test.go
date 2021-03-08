package dec04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBYR(t *testing.T) {
	assert.True(t, mandatoryKeys["byr"]("2002"))
	assert.False(t, mandatoryKeys["byr"]("2003"))
}

func TestHGT(t *testing.T) {
	assert.True(t, mandatoryKeys["hgt"]("60in"))
	assert.True(t, mandatoryKeys["hgt"]("190cm"))
	assert.False(t, mandatoryKeys["hgt"]("190in"))
	assert.False(t, mandatoryKeys["hgt"]("190"))
}

func TestHCL(t *testing.T) {
	assert.True(t, mandatoryKeys["hcl"]("#123abc"))
	assert.False(t, mandatoryKeys["hcl"]("#123abz"))
	assert.False(t, mandatoryKeys["hcl"]("123abd"))
}

func TestECL(t *testing.T) {
	assert.True(t, mandatoryKeys["ecl"]("brn"))
	assert.False(t, mandatoryKeys["ecl"]("wat"))
}

func TestPID(t *testing.T) {
	assert.True(t, mandatoryKeys["pid"]("000000001"))
	assert.False(t, mandatoryKeys["pid"]("0123456789"))
}
