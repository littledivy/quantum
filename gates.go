package main

// Identity Gate
func IdentityGate(width int) Gate {
	m := Identity(KetSize(width))
	return NewGate(width, m)
}

func PauliXGate() Gate {
	g := []complex128{complex(0, 0), complex(1, 0), complex(0, 0), complex(1, 0)}
	m := NewMatrixFromElements(2, g)
	return NewGate(1, m)
}
