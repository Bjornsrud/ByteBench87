package tests

import "github.com/Bjornsrud/ByteBench87/bench"

func DummyIntLoop() bench.TestCase {
	return bench.TestCase{
		Name: "DummyIntLoop",
		Run: func() uint64 {
			var x uint64 = 1
			for i := 0; i < 5_000_000; i++ {
				x = x*1664525 + 1013904223
			}
			return x
		},
	}
}
