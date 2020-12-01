package output

// Outputter defines what each exercise must implement.
type Outputter interface {
	Output() (string, error)
}
