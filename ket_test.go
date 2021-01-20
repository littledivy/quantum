package main

import (
	"log"
	"math"
	"testing"
)

func Try(err error, msg string) {
	if err != nil {
		log.Fatalf("error "+msg+": %v", err)
	}
}

func Assert(i interface{}) {
	if i == nil {
		log.Fatalf("Assertion failed: %v", i)
	}
}

func AssertEqual(i, j interface{}) {
	if i != j {
		log.Fatalf("Assertion failed: %v != %v", i, j)
	}
}

func TestValidKet(t *testing.T) {
	k := NewKet(3)
	k.Elements[0] = complex(0, 0)
	k.Elements[1] = complex(0, 0)
	k.Elements[2] = complex(1, 0)

	ki := NewKet(3)
	k.Elements[0] = complex(0.5, 0)
	k.Elements[1] = complex(0, 0.5)

	Assert(k.IsValid())
	AssertEqual(false, ki.IsValid())
}

func TestClassicalKet(t *testing.T) {
	sqrInv := 1 / math.Sqrt(2)

	cls := NewKet(3)
	cls.Elements[0] = complex(0, 0)
	cls.Elements[1] = complex(0, 0)
	cls.Elements[2] = complex(1, 0)
	Assert(cls.IsClassical())

	qt1 := NewKet(2)
	qt1.Elements[0] = complex(sqrInv, 0)
	qt1.Elements[1] = complex(0, sqrInv)
	AssertEqual(false, qt1.IsClassical())

	qt2 := NewKet(2)
	qt2.Elements[0] = complex(0.5, 0.5)
	qt2.Elements[1] = complex(0.5, 0.5)
	AssertEqual(false, qt2.IsClassical())
}
