package min_cost_climbing_stairs

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func MinCostClimbingStairs(cost []int) int {
	for position := len(cost) - 3; position >= 0; position-- {
		cost[position] += min(cost[position+1], cost[position+2])
	}

	return min(cost[0], cost[1])
}
