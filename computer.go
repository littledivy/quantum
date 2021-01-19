package main

type State int

const (
  Initializing State = iota
  Running
  Collapsed
)

type QuantumComputer struct {
  State State
  Width int
  Register QuantumRegister
  Classical ClassicalRegister
}

func NewQuantumComputer(width int) QuantumComputer {
  qr := NewQuantumRegister(width, ZeroedClassicalRegister(width))
  return QuantumComputer { State: Initializing, Width: width, Register: qr, Classical: ZeroedClassicalRegister(width) }
}

func (qc* QuantumComputer) Initialize(value int) {
  classical := NewClassicalRegisterFromState(qc.Width, value)
  qc.Register = NewQuantumRegister(qc.Width, classical)
  qc.State = Running
}

func (qc* QuantumComputer) Apply(gate Gate) {
  qc.Register.Apply(gate)
}

func (qc* QuantumComputer) Collapse() {
  qc.Classical = qc.Register.Collapse()
  qc.State = Collapsed
}

func (qc* QuantumComputer) Reset() {
  qc.State = Initializing
}

func (qc* QuantumComputer) Value() int {
  return qc.Classical.state()
}

func (qc* QuantumComputer) Probablities() []float64 {
  return qc.Register.Probablities()
}

