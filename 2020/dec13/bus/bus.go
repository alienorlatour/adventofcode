package bus

type Bus int

func New(id int) Bus {
	return Bus(id)
}

func (b Bus) FirstDeparture(t int64) int64 {
	bus := int64(b)
	if t%bus == 0 {
		// run !
		return t
	}
	return t + (bus - t%bus)
}

func (b Bus) ID() int {
	return int(b)
}
