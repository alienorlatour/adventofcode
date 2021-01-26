package docking

import (
	"strconv"
)

type Value interface {
	ApplyMask(m Mask) Value
	Uint64() uint64
}

type value uint64

func NewDockingValue(decimal string) (Value, error) {
	asInt, err := strconv.ParseUint(decimal, 10, 64)
	if err != nil {
		return nil, err
	}
	return value(asInt), nil
}

func (v value) ApplyMask(m Mask) Value {
	return (v | m.ones) & m.zeros
}

func (v value) Uint64() uint64 {
	return uint64(v)
}
