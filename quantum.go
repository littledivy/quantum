package main

import "fmt"

func main() {
	c1 := NewQuantumComputer(3)
	fmt.Println(c1.State == Initializing)
	c1.Initialize(5)
	c1.Apply(IdentityGate(3))
	fmt.Println(c1.State == Running)
	c1.Collapse()
	fmt.Println(c1.State == Collapsed)
	fmt.Println(c1.Value()) // 5
}
