package dec14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ablqk/adventofcode/2020/dec14/docking"
)

type Parser struct {
	// memory map[string]docking.Value // part 1 was so simple
	memory map[docking.Value]uint64
	mask   docking.Mask2
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
	p.mask, err = docking.NewMask2(s)
	return err
}

func (p *Parser) parseMemAssignation(s string) error {
	r := regexp.MustCompile("mem.([0-9]+). = ([0-9]+)")
	res := r.FindStringSubmatch(s)
	if len(res) != 3 {
		return fmt.Errorf("invalid line format: %s", s)
	}
	value, err := strconv.ParseUint(res[2], 10, 64)
	if err != nil {
		return err
	}
	key, err := docking.NewDockingValue(res[1])
	if err != nil {
		return err
	}

	for i := uint(0); i < 1<<p.mask.FloaterCount(); i++ {
		p.memory[p.mask.Floater(key, i)] = value
	}
	return nil
}

// Count adds up all the values
func (p *Parser) Count() uint64 {
	var count uint64
	for _, v := range p.memory {
		count += v
	}
	return count
}
