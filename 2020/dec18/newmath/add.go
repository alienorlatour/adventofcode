package newmath

type addition struct {
	left  expression
	right expression
}

func (a addition) compute() int {
	return a.left.compute() + a.right.compute()
}
