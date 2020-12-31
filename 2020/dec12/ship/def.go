package ship

type Function int

const (
	GoNorth Function = iota
	GoSouth
	GoEast
	GoWest
	TurnLeft
	TurnRight
	GoForward
)

type Instruction struct {
	Function Function
	Value    int
}

type Runner interface {
	Run() error
	Lattitude() int
	Longitude() int
}
