package compiler

import (
	"fmt"
	"strconv"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
	"github.com/ablqk/adventofcode/2020/dec12/ship/part2"
)

type Compiler struct {
	code []ship.Instruction
}

func (c *Compiler) ParseLine(s string) error {
	var instr ship.Instruction
	switch s[0] {
	case 'N':
		instr.Function = ship.North
	case 'S':
		instr.Function = ship.South
	case 'E':
		instr.Function = ship.East
	case 'W':
		instr.Function = ship.West
	case 'L':
		instr.Function = ship.Port
	case 'R':
		instr.Function = ship.Star
	case 'F':
		instr.Function = ship.Forward
	default:
		return fmt.Errorf("invalid instruction %s", s)
	}

	v, err := strconv.Atoi(s[1:])
	if err != nil {
		return fmt.Errorf("cannot parse value: %v", err)
	}
	instr.Value = v
	c.code = append(c.code, instr)
	return nil
}

func (c *Compiler) Compile() ship.Runner {
	// return part1.New(c.code, 90)
	return part2.New(c.code, 1, 10)
}
