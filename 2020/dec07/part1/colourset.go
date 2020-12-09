package part1

import (
	"github.com/ablqk/adventofcode/2020/dec07/defs"
)

type ColourSet struct {
	set map[defs.Colour]struct{}
}

func (cs *ColourSet) AppendAll(other *ColourSet) *ColourSet {
	for i := range other.set {
		cs.Append(i)
	}
	return cs
}

func (cs *ColourSet) Append(other defs.Colour) *ColourSet {
	if cs.set == nil {
		cs.set = make(map[defs.Colour]struct{})
	}
	cs.set[other] = struct{}{}
	return cs
}

func (cs *ColourSet) Slice() []defs.Colour {
	var s []defs.Colour
	for i := range cs.set {
		s = append(s, i)
	}
	return s
}
