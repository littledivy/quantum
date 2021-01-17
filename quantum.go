package main

import ("fmt"
 "math")

var R2 float32 = 0.70710678118

type QuantumGate struct {
  gate string
  qubit uint32
  target uint32
  angle float32
}

type QuantumCircuit struct {
  number_qubits uint32
  circuit []QuantumGate
}

func NewQuantumCircuit(q uint32) QuantumCircuit {
  return QuantumCircuit {number_qubits: q, circuit: []QuantumGate{}}
}

func (qc* QuantumCircuit) AddGate(gate string, qubit, target uint32, angle float32) {
  element :=  QuantumGate {gate, qubit, target, angle}
  qc.circuit = append(qc.circuit, element)
}

func (qc* QuantumCircuit) X(qubit uint32) {
  qc.AddGate("X", qubit, 0, 0.0)
}
func (qc* QuantumCircuit) Y(qubit uint32) {
  qc.RZ(qubit, math.Pi)
  qc.X(qubit)
}

func (qc* QuantumCircuit) H(qubit uint32) {
  qc.AddGate("H", qubit, 0, 0.0)
}

func (qc* QuantumCircuit) CX(qubit, target uint32) {
  qc.AddGate("CX", qubit, target, 0.0)
}

func (qc* QuantumCircuit) RX(qubit uint32, angle float32) {
  qc.AddGate("X", qubit, 0, angle)
}

func (qc* QuantumCircuit) RZ(qubit uint32, angle float32) {
  qc.H(qubit)
  qc.RX(qubit, angle)
  qc.H(qubit)
}

func (qc* QuantumCircuit) RY(qubit uint32, angle float32) {
  qc.RX(qubit, math.Pi/2.0)
  qc.H(qubit)
  qc.RX(qubit, angle)
  qc.H(qubit)
  qc.RX(qubit, -math.Pi/2.0)
}

type QuantumSimulator struct {
  circuit []QuantumGate
  number_qubits uint32
  bits uint32
  state_vector []complex128
}

func Simulator(circuit QuantumCircuit) QuantumSimulator {
  return QuantumSimulator { circuit: circuit.circuit, number_qubits: circuit.number_qubits, bits: circuit.number_qubits, state_vector: []complex128{} }
}

func (self* QuantumSimulator) init_state_vector() {
  for i := float64(1); i <= math.Pow(2, float64(self.number_qubits)); i++ {
    self.state_vector = append(self.state_vector, complex(0, 0))
    self.state_vector[0] = complex(1, 0)
  }
}

func turn(x, y complex128, angle float32) (complex128, complex128) {
  rx := float64(real(x))
  ry := float64(real(y))
  ix := float64(imag(x))
  iy := float64(imag(y))
  a := float64(angle / 2.0)
  a1 := complex(float64(R2) * (rx * math.Cos(a) + ry * math.Sin(a)), float64(R2) * (ix * math.Cos(a) + iy * math.Sin(a)))
  a2 := complex(float64(R2) * (rx * math.Cos(a) - ry * math.Sin(a)), float64(R2) * (ix * math.Cos(a) - iy * math.Sin(a)))
  return a1, a2
}

func superposition(x, y complex128) (complex128, complex128) {
  rx := float32(real(x))
  ry := float32(real(y))
  ix := float32(imag(x))
  iy := float32(imag(y))
  a1 := complex(R2 * (rx + ry), R2 * (ix + iy))
  a2 := complex(R2 * (rx - ry), R2 * (ix - iy))
  return complex128(a1), complex128(a2)
}

func (self* QuantumSimulator) Run() {
  self.init_state_vector()
  for _, quantum_gate := range self.circuit {
    if quantum_gate.gate == "X" {
    
    }
  }
}

func main() {
  qc := NewQuantumCircuit(15)
  qc.H(0)
  qc.CX(0,1)
  qc.X(1)
  fmt.Println("Executing circuit")
  qsim := Simulator(qc)
  qsim.Run()
}
