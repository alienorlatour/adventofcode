package defs

type Colour string

type ColourBags struct {
	Colour Colour
	Amount int
}

type BagRuleTree interface {
	ParseLine(s string) error
	Count(c Colour) int
}
