package dec01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ablqk/adventofcode/output"
)

const expectedTotal = 2020

func New(inputFile string) output.Outputter {
	return dec01{inputFile}
}

// dec01 is the implementation of this day's exercise
type dec01 struct {
	inputPath string
}

func (d dec01) Output() (string, error) {
	// read the input file
	lines, err := readLines(d.inputPath)
	if err != nil {
		return "", err
	}
	l, m, err := findMatch2(lines)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d * %d = %d", l, m, l*m), nil
}

func findMatch2(lines []int) (int, int, error) {
	// loop to find 2 values that add up to expectedTotal
	for i, l := range lines[:len(lines)-1] {
		for _, m := range lines[i+1:] {
			if l+m == expectedTotal {
				return l, m, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no match found")
}

// readLines reads a whole file into memory and returns a slice of its lines as ints
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// cat as int
		i, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}
