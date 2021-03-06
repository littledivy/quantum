package main

import "testing"

func TestClassicalRegister(t *testing.T) {
	n := NewClassicalRegister([]uint8{0, 1, 0, 1})
	if n.state() != 10 {
		t.Errorf("State = %d; want 10", n.state())
	}
}

func TestClassicalRegisterState(t *testing.T) {
	n := NewClassicalRegister([]uint8{0, 1, 0, 1})
	n2 := NewClassicalRegisterFromState(4, n.state())
	if n2.state() != 10 {
		t.Errorf("State = %d; want 10", n2.state())
	}
}
func TestClassicalRegister2(t *testing.T) {
	n := ZeroedClassicalRegister(4)
	r := NewQuantumRegister(4, n)
	end := r.Collapse()
	if n.state() != end.state() {
		t.Errorf("State = %d; want %d", n.state(), end.state())
	}
}
