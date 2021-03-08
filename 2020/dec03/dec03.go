package dec03

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

const (
	treeByte = '#' // representation of a tree
)

// New instance of the door for December 3rd
// 915 – Pope John X crowns Berengar I of Italy as Holy Roman Emperor.
func New(inputFile string, slopes [][2]int) doors.Solver {
	sl := make([]slope, len(slopes))
	for i, s := range slopes {
		sl[i] = slope{
			right: s[0],
			down:  s[1],
		}
	}
	return dec03{inputFile, sl}
}

// dec03 is the implementation of this day's exercise
type dec03 struct {
	inputPath string
	slopes    []slope
}

// slope represents the angle of the slope, in a way
type slope struct {
	right, down int
}

func (d dec03) Solve() (string, error) {
	// read the input file
	f, err := loadForest(d.inputPath)
	if err != nil {
		return "", err
	}

	result := 1 // will be multiplied by all the tree counts
	for _, sl := range d.slopes {
		// for each slope, multiply the number of plan by the previous result
		result *= f.countSlope(sl)
	}

	return fmt.Sprintf("Given slopes multiply to %d", result), nil
}

func loadForest(path string) (forest, error) {
	var f [][]byte
	err := fileread.ReadAndApply(path, func(s string) error {
		f = append(f, []byte(s))
		return nil
	})
	if err != nil {
		return forest{}, err
	}
	if len(f) == 0 {
		return forest{}, fmt.Errorf("no tree in the forest")
	}
	return forest{
		plan: f,
	}, nil
}
