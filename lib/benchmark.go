package lib

import (
	"log"
	"time"
)

func BenchmarkTime(function func()) {
	start := time.Now()
	defer func() {
		log.Printf("Benchmark execution time %s\n", time.Since(start))
	}()

	function()
}
