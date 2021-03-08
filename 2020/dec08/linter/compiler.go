package linter

import (
	"strconv"
)

type Compiler struct {
	code Code
}

// CompileLine adds the line to the code as something the machine can understand
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

// Lint finds the (first) infinite loop in the code and returns the accumulator value at that moment
func (comp *Compiler) Lint() (int, error) {
	return comp.code.Lint()
}

// Patch the code by trying each non-zero nop or any-value jmp and linting it until it works
func (comp *Compiler) Patch() int {
	accValue, err := comp.code.Lint()
	var fixLine int
	for err != nil && fixLine < len(comp.code.lines) {
		// look for next nop/jmp line
		for !canFix(comp.code.lines[fixLine]) {
			fixLine++
		}
		// fork the code, fix the line and lint the new version.
		f := comp.code.Fork()
		f.lines[fixLine] = fix(f.lines[fixLine])
		accValue, err = f.Lint()
		fixLine++ // don't try the same line in a loop
	}
	return accValue
}

// fix the line by changing the instruction
func fix(line codeLine) codeLine {
	switch line.instr {
	case jmp:
		line.instr = nop
	case nop:
		line.instr = jmp
	}
	return line
}

// canFix returns whether the line is a fix candidate.
func canFix(line codeLine) bool {
	return line.instr == jmp || // any jmp can be fixed
		(line.instr == nop && line.value != 0) // cannot fix nop into jmp 0
}
