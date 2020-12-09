package defs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseLine(s string, add func(container, contained Colour, amound int)) error {
	// extract the colour of the bag
	r := regexp.MustCompile("^([a-z]+ [a-z]+) bags contain (.*).")
	res := r.FindStringSubmatch(s)
	if len(res) != 3 {
		return fmt.Errorf("invalid line format, cannot parse container's colour: %s", s)
	}

	container := Colour(res[1]) // grape

	// check the special end-of-tree value
	if strings.Contains(res[2], "no other") {
		return nil
	}
	// extract each sub-bag
	r = regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+) bags?")
	ctnd := r.FindAllStringSubmatch(res[2], -1)
	// fill it up
	for _, exp := range ctnd {
		contained := Colour(exp[2])
		amount, err := strconv.Atoi(exp[1])
		if err != nil {
			return err
		}
		add(container, contained, amount)
	}
	return nil
}
