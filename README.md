## `quantum`

Advance quantum computer simuation in Go.

```golang
qc := NewQuantumComputer(3)
qc.Initialize(5)
qc.Apply(IdentityGate(3))
qc.Collapse()
fmt.Println(qc.Value())
```

### `related`

https://github.com/beneills/quantum/
https://github.com/qiskit-community/MicroQiskit


