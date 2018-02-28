package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestAll(t *testing.T) {
	a := fibDyn(0)
	b := fibRec(0)

	if a != b {
		spew.Dump(a, b)
		t.Error("u suk")
	}
}

func benchmark(b *testing.B, method, n int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch method {
		case 0:
			fibDyn(n)
		case 1:
			fibRec(n)
		default:
			panic("error")
		}
	}
}

func BenchmarkSeqMulti(b *testing.B) {
	b.Run("Seq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			seq(100000)
		}
	})
	b.Run("Multi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			multi(100000, 10)
		}
	})
}

func BenchmarkAll(b *testing.B) {
	methods := []string{"Dyn", "Rec"}
	x := 1
	for j := 0; j < 5; j++ {
		for i, m := range methods {
			b.Run(fmt.Sprint(m, x), func(b *testing.B) { benchmark(b, i, x) })
		}
		x *= 4
	}
}

func BenchmarkSmallDyn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibDyn(10)
	}
}
func BenchmarkSmallRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibRec(10)
	}
}
