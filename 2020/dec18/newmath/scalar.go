package newmath

type scalar struct {
	value int
}

func (a scalar) compute() int {
	return a.value
}
