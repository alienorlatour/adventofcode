package dec03

import (
	"fmt"

	"github.com/ablqk/adventofcode/doors"
	"github.com/ablqk/adventofcode/pkg/fileread"
)

// New instance of the Door for December 25
func New(input string) doors.Solver {
	return door{input}
}

type door struct {
	input string
}

// Solve the day's problem
func (d door) Solve() (string, error) {
	var count int

	var group []rucksack

	err := fileread.ReadAndApply(d.input, func(s string) error {
		// count += rucksack{s}.MistakePrio()

		group = append(group, rucksack{s})
		if len(group) == 3 {
			count += elfGroup{group}.Badge()
			group = nil
		}

		return nil
	})

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", count), nil
}

type elfGroup struct {
	elves []rucksack
}

func (g elfGroup) Badge() int {
	letter := g.findCommonLetter()
	return prio(letter)
}

func (g elfGroup) findCommonLetter() byte {
	mappedElves := make([]map[rune]struct{}, 3)
	for i, e := range g.elves {
		mappedElves[i] = countLetters(e.in)
	}

	for l := range mappedElves[0] {
		// yuuuuuck
		_, ok1 := mappedElves[1][l]
		_, ok2 := mappedElves[2][l]

		if ok1 && ok2 {
			return byte(l)
		}
	}

	return 0
}

type rucksack struct {
	in string
}

func (r rucksack) MistakePrio() int {
	pocketLeft, pocketRight := splitInTwo(r.in)
	commonLetter := findFirstCommonLetter(pocketLeft, pocketRight)
	return prio(commonLetter)
}

func prio(letter byte) int {
	if letter >= 'a' {
		return int(letter + 1 - 'a')
	}
	return int(letter + 27 - 'A')
}

func findFirstCommonLetter(left string, right string) byte {
	letters := countLetters(left)

	for _, r := range right {
		if _, ok := letters[r]; ok {
			return byte(r)
		}
	}

	return 0
}

func countLetters(left string) map[rune]struct{} {
	letters := make(map[rune]struct{})
	for _, r := range left {
		letters[r] = struct{}{}
	}
	return letters
}

func splitInTwo(in string) (string, string) {
	pocketSize := len(in) / 2
	return in[:pocketSize], in[pocketSize:]
}
