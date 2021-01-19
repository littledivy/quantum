package main

const MAX_SIZE int = 31

const MAX_ELEMENTS int = MAX_SIZE * MAX_SIZE

type Vector = []complex128

type Matrix struct {
  Size int
  Elements Vector
}

func NewMatrix(size int) Matrix {
  elements := make([]complex128, MAX_ELEMENTS)
  elements[0] = complex(0, 0)
  return Matrix { Size: size, Elements: elements }
}

func NewMatrixFromElements(size int, elements []complex128) Matrix {
  m := NewMatrix(size)
  for i, elem := range elements {
    m.Set(i / size, i % size, elem)
  }
  return m
}

func Identity(size int) Matrix {
  elements := make([]complex128, MAX_ELEMENTS)
  elements[0] = complex(0, 0)
  for i := 0; i <= size; i++ {
    elements[i * MAX_SIZE + i] = complex(1, 0)
  }
  
  return Matrix { Size: size, Elements: elements }
}

func (mat* Matrix) Embed(other Matrix, i, j int) {
  for x := 0; x <= other.Size; x++ {
    for y := 0; y <= other.Size; y++ {
      value :=  other.Get(x, y)
      mat.Set(i + x, i + y, value)
    }
  }
}

func (mat* Matrix) PermutateRows(permutation []int) Matrix {
  m := NewMatrix(mat.Size)
  for source_i, target_i := range permutation {
    for j := 0; j <= mat.Size; j++ {
      m.Set(target_i, j, mat.Get(source_i, j))
    }
  }
  
  return m
}

func (mat* Matrix) Set(i, j int, value complex128) {
  mat.Elements[i * MAX_SIZE + j] = value
}

func (mat* Matrix) Get(i, j int) complex128 {
  return mat.Elements[i * MAX_SIZE + j]
}

