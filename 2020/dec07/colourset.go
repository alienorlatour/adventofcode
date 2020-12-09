package dec07

type ColourSet struct {
	set map[Colour]struct{}
}

func (cs *ColourSet) AppendAll(other *ColourSet) *ColourSet {
	for i := range other.set {
		cs.Append(i)
	}
	return cs
}

func (cs *ColourSet) Append(other Colour) *ColourSet {
	if cs.set == nil {
		cs.set = make(map[Colour]struct{})
	}
	cs.set[other] = struct{}{}
	return cs
}

func (cs *ColourSet) Slice() []Colour {
	var s []Colour
	for i := range cs.set {
		s = append(s, i)
	}
	return s
}
