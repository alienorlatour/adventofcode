package docking

import (
	"strconv"
)

type Value uint64

func NewDockingValue(decimal string) (Value, error) {
	asInt, err := strconv.ParseUint(decimal, 10, 64)
	if err != nil {
		return 0, err
	}
	return Value(asInt), nil
}

func (v Value) Uint64() uint64 {
	return uint64(v)
}
