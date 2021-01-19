package main

type Gate struct {
  Width int
  Matrix Matrix
}

func NewGate(width int, matrix Matrix) Gate {
  return Gate { Width: width, Matrix: matrix }
}

func (gate* Gate) permutate(p []int) Gate {
  m := gate.Matrix.PermutateRows(p)
  return NewGate(gate.Width, m)
}
