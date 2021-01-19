package main

import (
	"math"
	"math/rand"
	"time"
)

type QuantumRegister struct {
	width     int
	collapsed bool
	ket       Ket
}

func NewQuantumRegister(width int, initial ClassicalRegister) QuantumRegister {
	return QuantumRegister{width: width, collapsed: false, ket: NewKetFromClassical(initial)}
}

func (qr *QuantumRegister) Apply(gate Gate) {
	qr.ket.Apply(gate)
}

func (qr *QuantumRegister) Probablities() []float64 {
	probablities := []float64{}
	kt := KetSize(qr.width)
	for _, coefficient := range qr.ket.Elements[kt:] {
		probablities = append(probablities, Csqr(coefficient))
	}

	return probablities
}

func (qr *QuantumRegister) Collapse() ClassicalRegister {
	rand.Seed(time.Now().UnixNano())
	qr.collapsed = true
	sample := math.Mod(rand.Float64(), 1)
	cumulative := 0.0

	for state, coefficient := range qr.ket.Elements {
		cumulative += Csqr(coefficient)
		if sample < cumulative {
			s := NewClassicalRegisterFromState(qr.width, int(state))
			return s
		}
	}
	return NewClassicalRegisterFromState(qr.width, 0)
}
