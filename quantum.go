package main

import (
	"fmt"
	"math"
	"reflect"
)

var R2 float32 = 0.70710678118

type QuantumGate struct {
	gate   string
	qubit  uint32
	target uint32
	angle  float32
}

type QuantumCircuit struct {
	number_qubits uint32
	circuit       []QuantumGate
}

func NewQuantumCircuit(q uint32) QuantumCircuit {
	return QuantumCircuit{number_qubits: q, circuit: []QuantumGate{}}
}

func (qc *QuantumCircuit) AddGate(gate string, qubit, target uint32, angle float32) {
	element := QuantumGate{gate, qubit, target, angle}
	qc.circuit = append(qc.circuit, element)
}

func (qc *QuantumCircuit) X(qubit uint32) {
	qc.AddGate("X", qubit, 0, 0.0)
}

func (qc *QuantumCircuit) Y(qubit uint32) {
	qc.RZ(qubit, math.Pi)
	qc.X(qubit)
}

func (qc *QuantumCircuit) Z(qubit uint32) {
	qc.RZ(qubit, math.Pi)
}

func (qc *QuantumCircuit) H(qubit uint32) {
	qc.AddGate("H", qubit, 0, 0.0)
}

func (qc *QuantumCircuit) CX(control, target uint32) {
	qc.AddGate("CX", control, target, 0.0)
}

func (qc *QuantumCircuit) RX(qubit uint32, angle float32) {
	qc.AddGate("RX", qubit, 0, angle)
}

func (qc *QuantumCircuit) RZ(qubit uint32, angle float32) {
	qc.H(qubit)
	qc.RX(qubit, angle)
	qc.H(qubit)
}

func (qc *QuantumCircuit) RY(qubit uint32, angle float32) {
	qc.RX(qubit, math.Pi/2.0)
	qc.H(qubit)
	qc.RX(qubit, angle)
	qc.H(qubit)
	qc.RX(qubit, -math.Pi/2.0)
}

type QuantumSimulator struct {
	circuit       []QuantumGate
	number_qubits uint32
	bits          uint32
	state_vector  []complex128
}

func Simulator(circuit QuantumCircuit) QuantumSimulator {
	return QuantumSimulator{circuit: circuit.circuit, number_qubits: circuit.number_qubits, bits: circuit.number_qubits, state_vector: []complex128{}}
}

func (self *QuantumSimulator) init_state_vector() {
	for i := float64(1); i <= math.Pow(2, float64(self.number_qubits)); i++ {
		self.state_vector = append(self.state_vector, complex(0, 0))
	}
	self.state_vector[0] = complex(1, 0)
}

func turn(x, y complex128, angle float32) (complex128, complex128) {
	rx := float64(real(x))
	ry := float64(real(y))
	ix := float64(imag(x))
	iy := float64(imag(y))
	a := float64(angle / 2.0)
	a1 := complex(rx*math.Cos(a)+iy*math.Sin(a), ix*math.Cos(a)-ry*math.Sin(a))
	a2 := complex(ry*math.Cos(a)+ix*math.Sin(a), iy*math.Cos(a)-rx*math.Sin(a))
	return a1, a2
}

func superposition(x, y complex128) (complex128, complex128) {
	rx := float32(real(x))
	ry := float32(real(y))
	ix := float32(imag(x))
	iy := float32(imag(y))
	a1 := complex(R2*(rx+ry), R2*(ix+iy))
	a2 := complex(R2*(rx-ry), R2*(ix-iy))
	return complex128(a1), complex128(a2)
}

func (self *QuantumSimulator) Run() {
	self.init_state_vector()
	for _, quantum_gate := range self.circuit {
		if quantum_gate.gate == "X" || quantum_gate.gate == "H" || quantum_gate.gate == "RX" {
			// Don't punish me typecasting gods
			for counter_qubit := uint32(0); counter_qubit <= uint32(math.Pow(2, float64(quantum_gate.qubit)))-1; counter_qubit++ {
				for counter_state := uint32(0); counter_state <= uint32(math.Pow(2, float64(self.number_qubits-quantum_gate.qubit-1)))-1; counter_state++ {
					qb0 := counter_qubit + uint32(math.Pow(2, float64(quantum_gate.qubit))+1)*counter_state
					qb1 := qb0 + uint32(math.Pow(2, float64(quantum_gate.qubit)))

					if quantum_gate.gate == "X" {
						swapF := reflect.Swapper(self.state_vector)
						swapF(int(qb0), int(qb1))
					}
					if quantum_gate.gate == "H" {
						a, b := superposition(self.state_vector[qb0], self.state_vector[qb1])
						self.state_vector[qb0] = a
						self.state_vector[qb1] = b
					}
					if quantum_gate.gate == "RX" {
						a, b := turn(self.state_vector[qb0], self.state_vector[qb1], quantum_gate.angle)
						self.state_vector[qb0] = a
						self.state_vector[qb1] = b
					}
				}
			}
		} else {
			low, high := lohi(quantum_gate.qubit, quantum_gate.target)
			swapF := reflect.Swapper(self.state_vector)
			// 0..2**low
			for cx0 := uint32(0); cx0 <= uint32(math.Pow(2, float64(low))); cx0++ {
			  // 2**(high-low-1)
				limit := math.Pow(2, float64(high-low-1))
				// 0..limit
				for cx1 := uint32(0); cx1 <= uint32(limit); cx1++ {
				  // 0..2**self.number_qubits - high - 1
					for cx2 := uint32(0); cx2 <= uint32(math.Pow(2, float64(self.number_qubits-high-1))); cx2++ {
					  // cx0 + 2**low+1 * cx1 + 2**high+1 * cx2 + 2**quantum_gate.qubit
						qb0 := cx0 + uint32(math.Pow(2, float64(low+1)))*cx1 + uint32(math.Pow(2, float64(high+1)))*cx2 + uint32(math.Pow(2, float64(quantum_gate.qubit)))
						// qb0 + 2**quantum_gate.target
						qb1 := qb0 + uint32(math.Pow(2, float64(quantum_gate.target)))
						swapF(int(qb0), int(qb1))
					}
				}
			}
		}
		self.print()
	}
}

func (self *QuantumSimulator) print() {
	for i := 1; i <= int(math.Pow(2, float64(self.number_qubits))-1); i++ {
		fmt.Printf("%04b %v, i%v\n", i, real(self.state_vector[i]), imag(self.state_vector[i]))
	}
}

func lohi(q0, q1 uint32) (uint32, uint32) {
	if q0 < q1 {
		return q0, q1
	} else {
		return q1, q0
	}
}

func main() {
	qc := NewQuantumCircuit(15)
	qc.H(0)
	for qubit := uint32(1); qubit <= 15; qubit++ {
		qc.H(qubit)
		qc.CX(qubit-1, qubit)
	}
	qc.H(0)
	qc.CX(0, 1)
	qc.X(1)
	qc.RX(2, math.Pi)
	qc.Z(0)
	qc.X(1)
	qc.RX(1, math.Pi)
	fmt.Println("Executing circuit")
	qsim := Simulator(qc)
	qsim.Run()
	fmt.Println(qc.number_qubits)
	fmt.Println(qc.circuit)
}
