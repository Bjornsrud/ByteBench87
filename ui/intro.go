package ui

import (
	"bufio"
	"fmt"
	"os"
)

const introPage1 = `
===========================================
  ByteBench87
  CPU Benchmarks • BYTE Magazine (1987)
===========================================

ByteBench87 is a hobby project inspired by the CPU benchmarks published in
BYTE Magazine in 1987, most notably the article "High-Tech Horsepower".

The original benchmarks compared the computational performance of CPUs
such as the Intel 80386 and the Motorola 68020 using small, focused tests.
They were designed to highlight differences in processor behavior rather
than to describe overall system performance.

This project reimplements those benchmarks in Go, aiming to recreate their
spirit rather than produce results directly comparable to the original data.

IMPORTANT:
This benchmark primarily measures single-core CPU performance and does NOT
represent overall system performance.

Press Enter to continue...
`

const introPage2 = `
-------------------------------------------
  Benchmark Execution Notes
-------------------------------------------

To improve consistency and reduce measurement noise, this benchmark
automatically performs:

• A warm-up phase before timing begins
• Multiple runs of each test
• Aggregation of results to produce stable measurements

For best results, it is recommended that you:

• Close all unnecessary applications
• Minimize background processes and system load
• Run the benchmark on an otherwise idle system
• Avoid heavy thermal or power throttling

Modern CPUs use features such as caching, frequency scaling, speculative
execution, and out-of-order processing. These can influence results and
introduce run-to-run variation, even under controlled conditions.

Results are intended for relative comparison and educational purposes only.

Press Enter to start benchmarks...
`

func waitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// ShowIntro displays the two-page introduction and disclaimer.
func ShowIntro() {
	fmt.Print(introPage1)
	waitForEnter()

	fmt.Print(introPage2)
	waitForEnter()
}
