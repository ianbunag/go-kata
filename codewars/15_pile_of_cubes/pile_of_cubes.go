package pile_of_cubes

// Average time complexity: O(n!)
// Worst time complexity:   O(n!)
// Space complexity:        O(1)
func FindNb(volume int) int {
	cubes := 1
	total := 0

	for {
		total = total + (cubes * cubes * cubes)

		if total == volume {
			return cubes
		}

		if total > volume {
			return -1
		}

		cubes += 1
	}
}
