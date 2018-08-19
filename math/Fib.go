package math

import "math"

type FibPred func (x uint64, ind uint64) bool

func FibWhile(aPred FibPred, bPred FibPred) [] uint64 {
	var (
		a uint64 = 0
		b uint64 = 1
		out = []uint64{}
		ind uint64 = 0
	)
	for aPred(a, ind) {
		out = append(out, a)
		ind += 1
		if bPred(b, ind) { out = append(out, b) }
		ind += 1
		a = a + b
		b = a + b
	}
	return out
}

func FibByLimit(limit uint64) [] uint64 {
	return FibWhile(
		func (a uint64, ind uint64) bool {
			return a <= limit
		},
		func (b uint64, ind uint64) bool {
			return b <= limit
		})
}

func FibByIterations(numInts uint64) [] uint64 {
	limit := numInts // uint64(math.Round(float64(numInts / 2)))
	lte := func (_, ind uint64) bool {
		return ind < limit
	}
	return FibWhile(lte, lte)
}

func NthFibNumber (nth uint64) uint64 {
	n := float64(nth)
	phi := math.Phi
	return uint64((math.Pow(phi, n) - math.Pow(-phi, -n)) / math.Sqrt(5.0))
}
