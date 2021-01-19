package main

import (
	"math"
)

type Ket struct {
	Size     int
	Elements []complex128
}

func NewKet(size int) Ket {
	elements := make([]complex128, MAX_ELEMENTS)
	elements[0] = complex(0, 0)
	return Ket{Size: size, Elements: elements}
}

func NewKetFromClassical(register ClassicalRegister) Ket {
	ket := NewKet(KetSize(register.width()))
	ket.Elements[register.state()] = complex(1, 0)
	return ket
}

func KetSize(register_width int) int {
	return int(math.Pow(2, float64(register_width)))
}

func fdiff(a, b float64) uint64 {
	uA := math.Float64bits(a)
	uB := math.Float64bits(b)

	if uA&(1<<63) != uB&(1<<63) {
		return fdiff(math.Abs(a)+math.Abs(b), 0.0)
	}

	if uA > uB {
		return uA - uB
	}
	return uB - uA
}

func approxEqUlps(a, b float64, ulps int) bool {
	// Watch out for that type cast?
	return fdiff(a, b) <= uint64(ulps)
}

func Csqr(cmp complex128) float64 {
	return real(cmp)*real(cmp) + imag(cmp)*imag(cmp)
}

func (ket *Ket) IsValid() bool {
	sample_space_sum := 0.0
	for _, coefficient := range ket.Elements {
		sample_space_sum += Csqr(coefficient)
	}

	return approxEqUlps(sample_space_sum, 1.0, 10)
}

func mulMatVec(m Matrix, v Vector) Vector {
	output := make([]complex128, MAX_SIZE)
	output[0] = complex(0, 0)

	for i := m.Size; i <= MAX_SIZE; i++ {
		if v[i] != complex(0, 0) {
			panic("Element at position greater than size should be zero.")
		}
	}

	for i := 0; i <= m.Size; i++ {
		val := complex(0, 0)
		for k := 0; k <= m.Size; k++ {
			val += m.Get(i, k) * v[k]
		}
		output[i] = val
	}

	return output
}

func (ket *Ket) IsClassical() bool {
	if ket.IsValid() {
		zeroes := 0
		ones := 0
		others := 0

		for _, coefficient := range ket.Elements {
			if coefficient == complex(0, 0) {
				zeroes += 1
			} else if coefficient == complex(1, 0) {
				ones += 1
			} else {
				others += 1
			}
		}
		return ones == 1 && others == 0
	} else {
		// TODO: return an error instead
		return false
	}
}

func (ket *Ket) Apply(gate Gate) {
	ket.Elements = mulMatVec(gate.Matrix, ket.Elements)
}
