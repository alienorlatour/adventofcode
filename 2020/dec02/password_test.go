package dec02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPasswordFromString(t *testing.T) {
	tt := map[string]struct {
		str      string
		expected password
	}{
		"nominal": {
			str: "1-3 a: abcde",
			expected: password{
				min:      1,
				max:      3,
				letter:   'a',
				password: []byte("abcde"),
			},
		},
		// todo add invalid case
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			p, err := newPasswordFromLine(tc.str)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, p)
		})
	}
}

func TestPassword_IsValidOldStyle(t *testing.T) {
	tt := map[string]struct {
		input    password
		expected bool
	}{
		"nominal": {
			input: password{
				min:      1,
				max:      3,
				letter:   'a',
				password: []byte("abcde"),
			},
			expected: true,
		},
		"invalid": {
			input: password{
				min:      1,
				max:      3,
				letter:   'b',
				password: []byte("cdefg"),
			},
			expected: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.input.IsValidOldStyle())
		})
	}
}
