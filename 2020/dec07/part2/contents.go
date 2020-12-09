package part2

import (
	"github.com/ablqk/adventofcode/2020/dec07/defs"
)

func NewBRT() defs.BagRuleTree {
	return &containedWiseBRT{}
}

// containedWiseBRT maps bags to the list of what they can contain
type containedWiseBRT struct {
	containers map[defs.Colour][]defs.ColourBags
}

func (b *containedWiseBRT) addContains(container defs.Colour, contained defs.Colour, amount int) {
	if b.containers == nil {
		b.containers = make(map[defs.Colour][]defs.ColourBags)
	}
	b.containers[container] = append(b.containers[container], defs.ColourBags{
		Colour: contained,
		Amount: amount,
	})
}

func (b *containedWiseBRT) ParseLine(s string) error {
	return defs.ParseLine(s, b.addContains)
}

// Count the total number of bags that will be contained in this specific colour
func (b *containedWiseBRT) Count(c defs.Colour) int {
	node, ok := b.containers[c]
	if !ok {
		// nothing
		return 0
	}
	var count int
	for _, v := range node {
		bagsInThisColour := b.Count(v.Colour) + 1
		count += bagsInThisColour * v.Amount
	}
	return count
}
