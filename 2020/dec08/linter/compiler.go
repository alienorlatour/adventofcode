package linter

import (
	"strconv"
)

type Compiler struct {
	code Code
}

func (comp *Compiler) CompileLine(s string) error {
	// instruction
	l := codeLine{
		instr: command(s[:3]),
	}
	// value
	v, err := strconv.Atoi(s[5:])
	if err != nil {
		return err
	}
	// sign
	if s[4] == '-' {
		v = -v
	}
	l.value = v
	comp.code.lines = append(comp.code.lines, l)
	return nil
}

func (comp *Compiler) Lint() int {
	return comp.code.Lint()
}
