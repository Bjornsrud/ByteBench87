package main

import (
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

	ui.PrintRunHeader(cfg)

	testCases := []bench.TestCase{
		// tests.DummyIntLoop(),
		tests.Fibonacci(),
		tests.Float(),
	}

	for _, tc := range testCases {
		ui.PrintTestHeader(tc.Name)
		result := bench.RunTest(tc, cfg)
		ui.PrintTestResult(result)
	}
}



