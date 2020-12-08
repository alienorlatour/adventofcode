package dec06

import (
	"fmt"

	"github.com/ablqk/adventofcode/2020/fileread"
	"github.com/ablqk/adventofcode/doors"
)

func New(input string) doors.Solver {
	return dec06{input}
}

type dec06 struct {
	input string
}

func (d dec06) Solve() (string, error) {
	var count int
	 g:= newGroup()
	err := fileread.ReadAndApply(d.input, func(s string) error {
		if s == "" {
			// this is the end of one group and the beginning of another
			count += g.CountYeses()
			g.Reset()
			return nil
		}
		// insert this answer
		g.InsertAnswer([]byte(s))
		return nil
	})
	count += g.CountYeses() // count the last group
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("The sum of all grouped yes answers is %d", count), nil
}
