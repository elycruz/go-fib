package main

import (
	"testing"
	"fmt"
	. "github.com/elycruz/go-fib/math"
)

func TestFibByIterations(t *testing.T) {
	t.Run("n={0-9}", func (t *testing.T) {
		// Cases map[n][]FibNumbs
		cases := make(map[uint64][]uint64)

		// Generate test case args
		for _, x := range uintRangeList(0, 9, 1) {
			cases[x] = fibByIters(x)
		}

		// Check ranges
		for k, _ := range cases {
			rslt := FibByIterations(k)
			//t.Log("Result:", rslt)
			//t.Log("Expected", cases[k])
			if !uintListEqual(cases[k], rslt) {
				t.Error()
			}
		}
	})
}

func TestFibByLimit(t *testing.T) {
	t.Run("limit={0-9}", func(t *testing.T) {
		// Cases map[n][]FibNumbs
		cases := make(map[uint64][]uint64)

		// Generate test case args
		for _, x := range uintRangeList(0, 9, 1) {
			cases[x] = fibByLimit(x)
		}

		// Check ranges
		for k, _ := range cases {
			rslt := FibByLimit(k)
			t.Log("Result:", rslt)
			t.Log("Expected", cases[k])
			if !uintListEqual(cases[k], rslt) {
				t.Error()
			}
		}
	})
}

func BenchmarkFibByIterations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range FibByIterations(100) {
			fmt.Print(n, " ")
		}
	}
}

func BenchmarkFibByLimit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range FibByLimit(9223372036854775808) { // hundreth fib num
			fmt.Print(n, " ")
		}
	}
}

func fibByIters (x uint64) []uint64 {
	var (
		a uint64 = 0
		b uint64 = 1
		out = []uint64{}
		ind uint64 = 0
	)
	for ind < x {
		out = append(out, a)
		ind += 1
		if ind < x { out = append(out, b) }
		ind += 1
		a = a + b
		b = a + b
	}
	return out
}

func fibByLimit (x uint64) []uint64 {
	var (
		a uint64 = 0
		b uint64 = 1
		out = []uint64{}
	)
	for a <= x {
		out = append(out, a)
		if b <= x { out = append(out, b) }
		a = a + b
		b = a + b
	}
	return out
}

func uintRangeList (min uint64, max uint64, step uint64) []uint64 {
	var (
		i = min
		out = []uint64{}
	)
	for ;i < max; i += step {
		out = append(out, i)
	}
	return out
}

func uintListEqual (xs1 []uint64, xs2 []uint64) bool {
	areEqual := len(xs1) == len(xs2)
	if !areEqual {
		return false;
	}
	for i, _ := range xs1 {
		if xs1[i] != xs2[i] {
			return false;
		}
	}
	return true
}
