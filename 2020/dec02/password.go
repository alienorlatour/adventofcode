package dec02

import (
	"fmt"
	"regexp"
	"strconv"
)

// newPasswordFromLine reads a line and parses the password's definition
func newPasswordFromLine(s string) (password, error) {
	r := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): (.*)")
	res := r.FindStringSubmatch(s)
	if len(res) != 5 {
		return password{}, fmt.Errorf("invalid line format: %s", s)
	}
	min, _ := strconv.Atoi(res[1])
	max, _ := strconv.Atoi(res[2])
	return password{
		min:      min,
		max:      max,
		letter:   res[3][0], // todo this is dangerous
		password: []byte(res[4]),
	}, nil
}

// password defines one line of the file.
// Consider defining a rule with min/max/letter and testing a string against it, rather than sticking them all together like some overcooked porridge.
type password struct {
	min, max int
	letter   byte
	password []byte
}

// IsValidOldStyle returns whether the password is valid according to the first part of the door
func (p password) IsValidOldStyle() bool {
	var count int
	for _, b := range p.password {
		if b == p.letter {
			count++
		}
	}
	return count <= p.max && count >= p.min
}

// IsValidNewStyle returns whether the password is valid according to the second part of the door
func (p password) IsValidNewStyle() bool {
	// we assume min < max
	if len(p.password) < p.max {
		// too short
		return false
	}
	// -1 because we start counting at 1
	minOK := p.password[p.min-1] == p.letter
	maxOK := p.password[p.max-1] == p.letter
	// This is a XOR
	return minOK != maxOK
}
