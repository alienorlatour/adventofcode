package dec18

import (
	"fmt"

	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 18
func New(input string) Dec18 {
	return Dec18{input}
}

type Dec18 struct {
	input string
}

// Solve the day's problem
func (d Dec18) Solve() (string, error) {
	var count int
	err := fileread.ReadAndApply(d.input, func(s string) error {
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

