package main

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
	"github.com/Bjornsrud/ByteBench87/tests"
	"github.com/Bjornsrud/ByteBench87/ui"
)

func main() {
	ui.ShowIntro()
	runBenchmarks()
}

func runBenchmarks() {
	cfg := bench.Config{WarmupRuns: 5, Runs: 11}

	fmt.Println("Starting benchmarks...")
	fmt.Printf("Config:   warmup=%d runs=%d\n\n", cfg.WarmupRuns, cfg.Runs)

	testCases := []bench.TestCase{
		tests.DummyIntLoop(),
		tests.Fibonacci(),
	}

	for _, tc := range testCases {
		fmt.Printf("Running %s...\n", tc.Name)

		result := bench.RunTest(tc, cfg)

		fmt.Println("Results:")
		printResult(result)
		fmt.Println()
	}
}


func printResult(r bench.Result) {
	fmt.Printf("Test:     %s\n", r.Name)
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n", r.Checksum)
}

