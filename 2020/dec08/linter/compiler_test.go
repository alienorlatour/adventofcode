package linter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompiler_CompileLine(t *testing.T) {
	tt := map[string]struct {
		line      string
		expected0 codeLine
	}{
		"acc": {
			"acc +5",
			codeLine{instr: acc, value: 5},
		},
		"jmp": {
			"jmp -264",
			codeLine{instr: jmp, value: -264},
		},
		// todo some invalid cases
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			compiler := Compiler{}
			err := compiler.CompileLine(tc.line)
			assert.NoError(t, err)
			assert.Equal(t, compiler.code.lines[0], tc.expected0)
		})
	}
}
