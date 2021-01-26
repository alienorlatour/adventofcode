package dec14

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ablqk/adventofcode/2020/dec14/docking"
)

type Parser struct {
	memory map[string]docking.Value
	mask   docking.Mask
}

func (p *Parser) Parse(s string) error {
	const maskPrefix = "mask = "
	if strings.HasPrefix(s, maskPrefix) {
		return p.parseNewMask(s[len(maskPrefix):])
	} else {
		// MemAssignation will check the validity of the line
		// no need to validate it here
		return p.parseMemAssignation(s)
	}
}

func (p *Parser) parseNewMask(s string) (err error) {
	p.mask, err = docking.NewMask(s)
	return err
}

func (p *Parser) parseMemAssignation(s string) error {
	r := regexp.MustCompile("mem.([0-9]+). = ([0-9]+)")
	res := r.FindStringSubmatch(s)
	if len(res) != 3 {
		return fmt.Errorf("invalid line format: %s", s)
	}
	value, err := docking.NewDockingValue(res[2])
	if err != nil {
		return err
	}
	p.memory[res[1]] = value.ApplyMask(p.mask)
	return nil
}

func (p *Parser) Count() uint64 {
	var count uint64
	for _, v := range p.memory {
		count += v.Uint64()
	}
	return count
}
