package dec05

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	colZero = "L"
	colOne  = "R"
	rowZero = "F"
	rowOne  = "B"
)

func NewBoardingPass(representation string) (BoardingPass, error) {
	r, c, err := parseBP(representation)
	if err != nil {
		return BoardingPass{}, err
	}
	return BoardingPass{r, c}, nil
}

// parseBP returns the row and column of a string BoardingPass
func parseBP(repr string) (int, int, error) {
	if len(repr) != 10 {
		return 0, 0, fmt.Errorf("invalid length, should be 10, got %d", len(repr))
	}
	r, err := parseBinary(repr[:7], rowZero, rowOne)
	if err != nil {
		return 0, 0, err
	}
	c, err := parseBinary(repr[7:], colZero, colOne)
	if err != nil {
		return 0, 0, err
	}
	return int(r), int(c), nil
}

func parseBinary(col string, zero, one string) (uint64, error) {
	col = strings.ReplaceAll(col, zero, "0")
	col = strings.ReplaceAll(col, one, "1")
	return strconv.ParseUint(col, 2, 64)
}

type BoardingPass struct {
	row    int
	column int
}

// IsHigher returns whether this BoardingPass is strictly higher than another
func (bp BoardingPass) IsHigher(other BoardingPass) bool {
	return bp.SeatID() > other.SeatID()
}

func (bp BoardingPass) SeatID() int {
	return bp.row*8 + bp.column
}
