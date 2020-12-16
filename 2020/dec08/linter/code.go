package linter

import (
	"fmt"
)

type instruction int

const (
	jmp instruction = iota
	acc
	nop
)

func command(s string) instruction {
	switch s {
	case "jmp":
		return jmp
	case "acc":
		return acc
	default:
		// for the sake of fluidity, anything else is a nop
		return nop
	}
}

type codeLine struct {
	instr instruction
	value int
	// in a better world, a line has no reason to know whether it has been visited or not
	// this should reside elsewhere
	visited bool
}

type Code struct {
	lines    []codeLine
	accum    int
	execLine int
}

// Lint returns the value of the accumulator when it encounters an infinite loop, -1 if none
func (c *Code) Lint() (int, error) {
	if c.execLine == len(c.lines) {
		return c.accum, nil // perfect !
	}
	if c.execLine < 0 || c.execLine > len(c.lines) {
		return -1, fmt.Errorf("the code jumped out of itself, cannot execute line %d of %d", c.execLine, len(c.lines))
	}
	if c.lines[c.execLine].visited {
		return c.accum, fmt.Errorf("infinite loop found on line %d", c.execLine)
	}
	c.lines[c.execLine].visited = true // linter
	// execute the line
	c.execute()
	return c.Lint()
}

func (c *Code) execute() {
	line := c.lines[c.execLine]
	switch line.instr {
	case jmp:
		c.execLine += line.value
	case acc:
		c.accum += line.value
		c.execLine++
	case nop:
		c.execLine++
	}
}

func (c *Code) Fork() *Code {
	// copy array
	newCode := make([]codeLine, len(c.lines))
	for i, l := range c.lines {
		// copy only instruction and value
		newCode[i] = codeLine{instr: l.instr, value: l.value}
	}
	return &Code{lines: newCode}
}
