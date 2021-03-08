package part1

import (
	"github.com/ablqk/adventofcode/2020/dec07/defs"
)

func NewBRT() defs.BagRuleTree {
	return &containersWiseBRT{}
}

// containersWiseBRT maps bags to the list of what can contain them
type containersWiseBRT struct {
	containers map[defs.Colour][]defs.Colour
}

func (b *containersWiseBRT) addContains(container defs.Colour, contained defs.Colour, amount int) {
	if b.containers == nil {
		b.containers = make(map[defs.Colour][]defs.Colour)
	}
	b.containers[contained] = append(b.containers[contained], container)
}

func (b *containersWiseBRT) Count(c defs.Colour) int {
	return len(b.allContainers(c).Slice())
}

func (b *containersWiseBRT) allContainers(c defs.Colour) *ColourSet {
	containers := &ColourSet{}
	node, ok := b.containers[c]
	if !ok {
		// this node has no possible container
		return containers
	}
	for _, i := range node {
		// append the whole tree of containers
		containers.Append(i)
		containers = containers.AppendAll(b.allContainers(i))
	}
	return containers
}

func (b *containersWiseBRT) ParseLine(s string) error {
	return defs.ParseLine(s, b.addContains)
}
