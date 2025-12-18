package bench

import (
	"sort"
	"time"
)

// globalSink is used to ensure benchmark computations are observable
// outside the benchmarked function, preventing dead-code elimination.
var globalSink uint64

// TestCase represents a single benchmark test.
type TestCase struct {
	Name string
	Run  func() uint64
}

// Result contains timing data for one test.
type Result struct {
	Name     string
	Times    []time.Duration
	Median   time.Duration
	Checksum uint64
}

// Config controls how benchmarks are executed.
type Config struct {
	WarmupRuns int
	Runs       int
}

func RunTest(tc TestCase, cfg Config) Result {
	// Warm-up (not timed)
	for i := 0; i < cfg.WarmupRuns; i++ {
		globalSink ^= mix64(tc.Run() + uint64(i))
	}

	times := make([]time.Duration, 0, cfg.Runs)

	var checksum uint64 = 0xcbf29ce484222325 // seed (FNV-ish)
	for i := 0; i < cfg.Runs; i++ {
		start := time.Now()
		v := tc.Run()
		elapsed := time.Since(start)

		times = append(times, elapsed)

		// Fold the computation result into a checksum so the compiler
		// cannot treat the work as unused.
		checksum ^= mix64(v + uint64(i))
		checksum *= 0x100000001b3
	}

	median := computeMedian(times)

	// Make it observable outside this function as well.
	globalSink ^= checksum

	return Result{
		Name:     tc.Name,
		Times:    times,
		Median:   median,
		Checksum: checksum,
	}
}

func computeMedian(times []time.Duration) time.Duration {
	if len(times) == 0 {
		return 0
	}

	sorted := append([]time.Duration(nil), times...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func mix64(x uint64) uint64 {
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33
	return x
}
