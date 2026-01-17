package ui

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
)

const seperator = "----------------------------"

func PrintRunHeader(cfg bench.Config) {
	fmt.Println("Starting benchmarks...")
	fmt.Printf("Config:   warmup=%d runs=%d\n\n", cfg.WarmupRuns, cfg.Runs)
}

func PrintTestHeader(name string) {
	fmt.Printf("\nRunning %s...\n", name)
}

func PrintTestResult(r bench.Result) {
	fmt.Printf(seperator + "\n")
	fmt.Printf("Test:     %s\n", r.Name)
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n", r.Checksum)
	fmt.Printf(seperator + "\n")
}
