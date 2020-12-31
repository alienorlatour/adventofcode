package compiler

import (
	"fmt"
	"strconv"

	"github.com/ablqk/adventofcode/2020/dec12/ship"
	"github.com/ablqk/adventofcode/2020/dec12/ship/part1"
)

type Compiler struct {
	code []ship.Instruction
}

func (c *Compiler) ParseLine(s string) error {
	var instr ship.Instruction
	switch s[0] {
	case 'N':
		instr.Function = ship.GoNorth
	case 'S':
		instr.Function = ship.GoSouth
	case 'E':
		instr.Function = ship.GoEast
	case 'W':
		instr.Function = ship.GoWest
	case 'L':
		instr.Function = ship.TurnLeft
	case 'R':
		instr.Function = ship.TurnRight
	case 'F':
		instr.Function = ship.GoForward
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
	return part1.New(c.code, 90)
}
