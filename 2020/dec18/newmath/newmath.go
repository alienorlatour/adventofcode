package newmath

// Compute the result of an expression given as a string
func Compute(in string) (int, error) {
	expr, err := parse(in)
	if err != nil {
		return 0, err
	}
	return expr.compute(), nil
}

type expression interface {
	compute() int
}
