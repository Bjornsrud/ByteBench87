package tests

import "github.com/Bjornsrud/ByteBench87/bench"

// Float is a benchmark test based on C code published by BYTE Magazine (July 1987).
// It stresses floating-point arithmetic using repeated multiplications
// and divisions in a tight loop, similar in spirit to Listing 2.
func Float() bench.TestCase {
	return bench.TestCase{
		Name: "Float",
		Run: func() uint64 {
			const count = 31_000_000

			a := 1.23456789
			b := 3.14159265
			c := 0.0

			for i := 0; i < count; i++ {
				// Make each iteration depend (slightly) on i to prevent
				// the compiler from collapsing the loop.
				t := float64((i & 1023) + 1) // 1..1024 repeating

				c = (a*b + t) / (b + 0.000001*t)
				a = (c*b + 0.000001*t) / (b - 0.000001*t)
			}

			// Mix both a and c into the return value so it depends on the loop.
			// Also fold in 'count' so changes to workload are reflected in the checksum.
			ai := uint64(a * 1e9)
			ci := uint64(c * 1e9)
			return (ai ^ (ci << 1)) ^ uint64(count)
		},
	}
}

