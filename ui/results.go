package ui

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
)

const separator = "----------------------------"

func PrintRunHeader(cfg bench.Config) {
	fmt.Println("Starting benchmarks...")
	fmt.Printf("Config:   warmup=%d runs=%d\n\n", cfg.WarmupRuns, cfg.Runs)
}

func PrintTestHeader(name string) {
	fmt.Printf("\nRunning %s...\n", name)
}

func PrintTestResult(r bench.Result) {
	fmt.Printf(separator + "\n")
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n", r.Checksum)
	fmt.Printf(separator + "\n")
}
