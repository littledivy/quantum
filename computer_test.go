package main

import "testing"

func TestNewComputer(t *testing.T) {
	c1 := NewQuantumComputer(3)
	if c1.State != Initializing {
		t.Errorf("Assertion failed: Computer state should be Initializing.")
	}
}

func TestIdentityGate(t *testing.T) {
	c1 := NewQuantumComputer(3)
	if c1.State != Initializing {
		t.Errorf("Assertion failed: Computer state should be Initializing.")
	}

	c1.Initialize(5)
	c1.Apply(IdentityGate(3))
	if c1.State != Running {
		t.Errorf("Assertion failed: Computer state should be Running.")
	}

	c1.Collapse()
	if c1.State != Collapsed {
		t.Errorf("Assertion failed: Computer state should be Collapsed.")
	}

	if c1.Value() != 5 {
		t.Errorf("Assertion failed: Identity Gate computation failed.")
	}
}
