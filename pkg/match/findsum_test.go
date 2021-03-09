package match

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindMatch2(t *testing.T) {
	tt := map[string]struct {
		haystack       []int
		needle         int
		expectedErr    string
		expectedValues [2]int
	}{
		"original example": {
			haystack:       []int{1721, 979, 366, 299, 675, 1456},
			expectedValues: [2]int{1721, 299},
			needle:         2020,
		},
		"no match": {
			haystack:    []int{1, 2, 3, 4, 5, 6},
			expectedErr: "no match found",
			needle:      202,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			l, m, err := FindSum2(tc.haystack, tc.needle)
			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err.Error())
			}
			assert.Equal(t, tc.expectedValues[0], l)
			assert.Equal(t, tc.expectedValues[1], m)
		})
	}
}

func benchFindSum2_Size(lines []int, expected int, b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		_, _, _ = FindSum2old(lines, 2020)
	})
	b.Run("newer", func(b *testing.B) {
		_, _, _ = FindSum2(lines, 2020)
	})
}

func BenchmarkFindSum2_32(b *testing.B) {
	lines := randomInts(32)
	lines[16] = 1729
	lines[30] = 291
	for n := 0; n < b.N; n++ {
		benchFindSum2_Size(lines, 2020, b)
	}
}

func BenchmarkFindSum2_10k(b *testing.B) {
	lines := randomInts(10000)
	lines[5000] = 1729
	lines[9899] = 291
	for n := 0; n < b.N; n++ {
		benchFindSum2_Size(lines, 2020, b)
	}
}

func BenchmarkFindSum2_100k(b *testing.B) {
	lines := randomInts(100000)
	lines[16023] = 1729
	lines[30971] = 291
	for n := 0; n < b.N; n++ {
		benchFindSum2_Size(lines, 2020, b)
	}
}

func randomInts(size int) []int {
	lines := make([]int, size)
	for i := 0; i < size; i++ {
		lines[i] = rand.Int() & 1023 + 2100 // no 2 values will add up to 2020
	}
	return lines
}
