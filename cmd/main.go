package main

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
	"github.com/Bjornsrud/ByteBench87/ui"
)

func main() {
	ui.ShowIntro()
	runBenchmarks()
}

func runBenchmarks() {
	cfg := bench.Config{
		WarmupRuns: 3,
		Runs:       5,
	}

	test := bench.TestCase{
		Name: "DummyIntLoop",
		Run: func() uint64 {
			var x uint64 = 1
			for i := 0; i < 5_000_000; i++ {
				x = x*1664525 + 1013904223
			}
			return x
		},
	}

	result := bench.RunTest(test, cfg)

	printResult(result)
}

func printResult(r bench.Result) {
	fmt.Printf("Test:     %s\n", r.Name)
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n", r.Checksum)
}

