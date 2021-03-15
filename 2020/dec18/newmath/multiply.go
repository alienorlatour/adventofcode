package newmath


type multiplication struct {
	left  expression
	right expression
}

func (a multiplication) compute() int {
	return a.left.compute() * a.right.compute()
}

