package main

import (
	"math"
)

type ClassicalRegister struct {
	bits []uint8
}

func NewClassicalRegister(bits []uint8) ClassicalRegister {
	return ClassicalRegister{bits: bits}
}

func NewClassicalRegisterFromState(width, state int) ClassicalRegister {
	bits := []uint8{}
	remaining_state := state
	for i := 0; i <= width; i++ {
		pos := width - i - 1
		value := int(math.Pow(2, float64(pos)))
		if value <= remaining_state {
			remaining_state -= value
			bits = append([]uint8{1}, bits...)
		} else {
			bits = append([]uint8{0}, bits...)
		}
	}
	bits = bits[1:]
	return NewClassicalRegister(bits)
}

func (cg *ClassicalRegister) width() int {
	return len(cg.bits)
}

func ZeroedClassicalRegister(width int) ClassicalRegister {
	bits := make([]uint8, width)
	bits[0] = 0
	return NewClassicalRegister(bits)
}

func (cg *ClassicalRegister) state() int {
	state := 0
	for pos, bit := range cg.bits {
		if bit != 0 {
			state += int(math.Pow(2, float64(pos)))
		}
	}
	return state
}
