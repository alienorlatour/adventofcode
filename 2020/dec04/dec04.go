package dec04

import (
	"fmt"
	"strings"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/libs/fileread"
)

const (
	entrySeparator = " "
	kvSeparator    = ":"
)

// New instance of the Door for December 4
// 1259 – Kings Louis IX of France and Henry III of England agree to the Treaty of Paris, in which Henry renounces his claims to French-controlled territory on continental Europe (including Normandy) in exchange for Louis withdrawing his support for English rebels.
func New(input string) doors.Solver {
	return dec04{input}
}

type dec04 struct {
	input string
}

// Solve the day's door
func (d dec04) Solve() (string, error) {
	// read the input file
	var valid int
	p := newPassport()
	err := fileread.ReadAndApply(d.input, func(s string) error {
		if s == "" {
			// this is the end of one passport and the beginning of another
			if p.IsValid() {
				valid++
			}
			p.Reset()
			return nil
		}

		// do the splitting and populate the passport
		entries := strings.Split(s, entrySeparator)
		for _, e := range entries {
			kv := strings.Split(e, kvSeparator)
			if len(kv) != 2 {
				return fmt.Errorf("invalid entry %s", e)
			}
			p.Add(kv[0], kv[1])
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Found %d valid passports", valid), nil
}
