package lib_benchmark

import (
	"log"
	"time"
)

func BenchmarkTime(description string, function func()) {
	start := time.Now()
	defer func() {
		log.Printf("%s benchmark execution time %s\n", description, time.Since(start))
	}()

	function()
}
