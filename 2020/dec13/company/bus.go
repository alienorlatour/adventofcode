package company

type Bus int

func (b Bus) firstDeparture(t int) int {
	bus := int(b)
	if t%bus == 0 {
		// run !
		return t
	}
	return t + (bus - t%bus)
}
