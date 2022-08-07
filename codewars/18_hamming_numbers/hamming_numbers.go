package hamming_numbers

import (
	"sync"

	"github.com/yvnbunag/go-kata/lib"
)

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func Hammer(n int) uint {
	return 0
}

func CalculateHammingNumber(i, j, k uint) uint {
	var result uint = 1
	resultsChannel := make(chan uint, 3)
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib.PowerUInt(2, i)
	}()
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib.PowerUInt(3, j)
	}()
	go func() {
		defer waitGroup.Done()
		resultsChannel <- lib.PowerUInt(5, k)
	}()
	go func() {
		waitGroup.Wait()
		close(resultsChannel)
	}()

	for value := range resultsChannel {
		result *= value
	}

	return result
}
