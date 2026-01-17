package ui

import (
	"fmt"

	"github.com/Bjornsrud/ByteBench87/bench"
)

func PrintRunHeader(cfg bench.Config) {
	fmt.Println("Starting benchmarks...")
	fmt.Printf("Config:   warmup=%d runs=%d\n\n", cfg.WarmupRuns, cfg.Runs)
}

func PrintTestHeader(name string) {
	fmt.Printf("Running %s...\n", name)
}

func PrintTestResult(r bench.Result) {
	fmt.Println("*** Summary ***")
	fmt.Printf("Test:     %s\n", r.Name)
	fmt.Printf("Median:   %v\n", r.Median)
	fmt.Printf("Checksum: 0x%016x\n\n", r.Checksum)
}
