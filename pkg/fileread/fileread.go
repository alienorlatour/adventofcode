package fileread

import (
	"bufio"
	"os"
	"strconv"
)

// ReadAndApply reads a file and applies a func on each line
func ReadAndApply(path string, f func(s string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = f(scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

// ReadStrings returns a slice of all the lines in a file
func ReadStrings(path string) ([]string, error) {
	var lines []string
	err := ReadAndApply(path, func(s string) error {
		lines = append(lines, s)
		return nil
	})
	return lines, err
}

// ReadInts returns a slice of all lines in a file, as ints
func ReadInts(path string) ([]int, error) {
	var lines []int
	err := ReadAndApply(path, func(s string) error {
		// cat as int
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		lines = append(lines, i)
		return nil
	})
	return lines, err
}
