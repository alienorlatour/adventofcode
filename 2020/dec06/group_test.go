package dec06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupPartOne(t *testing.T) {
	tt := map[string]struct {
		answers       []string
		expectedCount int
	}{
		"abc": {
			answers:       []string{"abc", "ab", "b"},
			expectedCount: 3,
		},
		"one answer": {
			answers:       []string{"abc"},
			expectedCount: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := newGroupPartOne()
			for _, s := range tc.answers {
				g.InsertAnswer([]byte(s))
			}
			assert.Equal(t, tc.expectedCount, g.CountYeses())
		})
	}
}

func TestGroupPartTwo(t *testing.T) {
	tt := map[string]struct {
		answers       []string
		expectedCount int
	}{
		"b": {
			answers:       []string{"abc", "ab", "banana", "byte"},
			expectedCount: 1,
		},
		"one answer": {
			answers:       []string{"abc"},
			expectedCount: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := newGroupPartTwo()
			for _, s := range tc.answers {
				g.InsertAnswer([]byte(s))
			}
			assert.Equal(t, tc.expectedCount, g.CountYeses())
		})
	}
}
