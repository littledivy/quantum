package main

import "math"

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

func (qc QuantumCircuit) AddGate(gate string, qubit, target uint32, angle float32) {
  element :=  QuantumGate {gate, qubit, target, angle}
  qc.circuit = append(qc.circuit, element)
}

func (qc QuantumCircuit) X(qubit uint32) {
  qc.AddGate("X", qubit, 0, 0.0)
}
func (qc QuantumCircuit) Y(qubit uint32) {
  qc.RZ(qubit, math.Pi)
  qc.X(qubit)
}

func (qc QuantumCircuit) H(qubit uint32) {
  qc.AddGate("H", qubit, 0, 0.0)
}

func (qc QuantumCircuit) CX(qubit, target uint32) {
  qc.AddGate("CX", qubit, target, 0.0)
}

func (qc QuantumCircuit) RX(qubit uint32, angle float32) {
  qc.AddGate("X", qubit, 0, angle)
}

func (qc QuantumCircuit) RZ(qubit uint32, angle float32) {
  qc.H(qubit)
  qc.RX(qubit, angle)
  qc.H(qubit)
}

func (qc QuantumCircuit) RY(qubit uint32, angle float32) {
  qc.RX(qubit, math.Pi/2.0)
  qc.H(qubit)
  qc.RX(qubit, angle)
  qc.H(qubit)
  qc.RX(qubit, -math.Pi/2.0)
}

func main() {}
