package tests

import "github.com/Bjornsrud/ByteBench87/bench"

// Fibonacci returns a benchmark test inspired by BYTE Magazine (1987).
// It uses a naive recursive implementation to stress function calls
// and integer arithmetic.
func Fibonacci() bench.TestCase {
	return bench.TestCase{
		Name: "Fibonacci",
		Run: func() uint64 {
			const n = 33
			const reps = 20

			var acc uint64
			for i := 0; i < reps; i++ {
				acc ^= fib(n) // or change to fib(n + (i & 1)) to add tiny variation to avoid overly predictable paths
			}
			return acc
		},
	}
}

func fib(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	return fib(n-1) + fib(n-2)
}
