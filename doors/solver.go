package doors

// Solver defines what each calendar door must implement.
type Solver interface {
	Solve() (string, error)
}
