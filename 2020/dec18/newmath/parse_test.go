package newmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parser_parseNumSymb(t *testing.T) {
	tt := map[string]struct {
		expected expression
		err      string
	}{
		"3": {
			expected: scalar{3}},
		"3+2": {
			expected: addition{left: scalar{3}},
		},
		"6598*2": {
			expected: multiplication{left: scalar{6598}},
		},
		"letters": {
			err: "invalid number",
		},
		"34invalid": {
			err: "invalid symbol",
		},
	}
	for in, tc := range tt {
		t.Run(in, func(t *testing.T) {
			p := parser{input: in}
			p.parseNumSymb()
			if tc.err == "" {
				assert.NoError(t, p.err)
				assert.Equal(t, tc.expected, p.tree)
			} else {
				assert.Error(t, p.err)
				assert.Contains(t, p.err.Error(), tc.err)
			}
		})
	}
}

func Test_parser_parseParens(t *testing.T) {
	tt := map[string]struct {
		expected      expression
		err           string
		expectedIndex int
	}{
		"(3)": {
			expected: scalar{3},
			expectedIndex: 3,
		},
		"(3+2)": {
			expected: addition{left: scalar{3}, right: scalar{2}},
			expectedIndex: 5,
		},
		"(6598*2)+3": {
			expected: multiplication{left: scalar{6598}},
			expectedIndex: 8,
		},
		"((3+2)+((5+3)+3)+1)+5": {
			expectedIndex: 17,
			expected: addition{left: addition{left: addition{left: scalar{3}, right: scalar{2}}}},
		},
	}
	for input, tc := range tt {
		t.Run(input, func(t *testing.T) {
			p := parser{input: input}
			p.parseParens()
			if tc.err == "" {
				assert.NoError(t, p.err)
				assert.Equal(t, tc.expected, p.tree)
			} else {
				assert.Error(t, p.err)
				assert.Contains(t, p.err.Error(), tc.err)
			}
		})
	}
}
