package dec10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	in := []int{
		1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19,
	}
	assert.Equal(t, []int{3, 2, 1}, compress(in))
}

func TestArrangements(t *testing.T) {
	assert.Equal(t, 1, arrangements(0))
	assert.Equal(t, 1, arrangements(1))
	assert.Equal(t, 2, arrangements(2))
	assert.Equal(t, 4, arrangements(3))
	assert.Equal(t, 7, arrangements(4))
}
