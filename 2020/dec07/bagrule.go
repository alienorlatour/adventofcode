package dec07

import (
	"fmt"
	"regexp"
	"strings"
)

type Colour string

type BagRuleTree struct {
	// containers mapps bags to the list of what can contain them
	containers map[Colour][]Colour
}

func (b *BagRuleTree) AddContains(container Colour, contained Colour) {
	if b.containers == nil {
		b.containers = make(map[Colour][]Colour)
	}
	b.containers[contained] = append(b.containers[contained], container)
}

func (b *BagRuleTree) parseLine(s string) error {
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
		b.AddContains(container, contained)
	}
	return nil
}

func (b *BagRuleTree) Containers(c Colour) *ColourSet {
	containers := &ColourSet{}
	node, ok := b.containers[c]
	if !ok {
		// this node has no possible container
		return containers
	}
	for _, i := range node {
		// append the whole tree of containers
		containers.Append(i)
		containers = containers.AppendAll(b.Containers(i))
	}
	return containers
}
