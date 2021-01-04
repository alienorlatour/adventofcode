package ship

type Function int

const (
	North Function = iota
	South
	East
	West
	Port
	Star
	Forward
)

type Instruction struct {
	Function Function
	Value    int
}

type Runner interface {
	Run() error
	Latitude() int
	Longitude() int
}
