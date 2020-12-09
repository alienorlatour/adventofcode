package linter

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
func (c *Code) Lint() int {
	if c.execLine >= len(c.lines) {
		return -1 // strange way to exit
	}
	if c.lines[c.execLine].visited {
		return c.accum
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
