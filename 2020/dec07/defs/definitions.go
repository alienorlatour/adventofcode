package defs

// Colour defines a bag colour
type Colour string

// ColourBags defines a set of bags of one single colour
type ColourBags struct {
	Colour Colour
	Amount int
}

// BagRuleTree is a tree of rules for bags
type BagRuleTree interface {
	ParseLine(s string) error
	Count(c Colour) int
}
