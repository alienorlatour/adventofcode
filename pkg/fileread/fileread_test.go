package fileread

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStrings(t *testing.T) {
	s, err := ReadStrings("./testdata/in.txt")
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "5", "6", "9", "4", "7836", "4893", "5125", "3484", "2546", "25", "27"}, s)
}

func TestReadInts(t *testing.T) {
	s, err := ReadInts("./testdata/in.txt")
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 5, 6, 9, 4, 7836, 4893, 5125, 3484, 2546, 25, 27}, s)
}
