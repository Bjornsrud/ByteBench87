package main

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
	"github.com/Bjornsrud/ByteBench87/ui"
	"github.com/Bjornsrud/ByteBench87/tests"
	
)

func main() {
	ui.ShowIntro()
	runBenchmarks()
}

func runBenchmarks() {
	cfg := bench.Config{WarmupRuns: 3, Runs: 5}
	result := bench.RunTest(tests.DummyIntLoop(), cfg)
	printResult(result)
}

func printResult(r bench.Result) {
	fmt.Printf("Test:     %s\n", r.Name)
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n", r.Checksum)
}
