package beeramid

// Average time complexity: O(n!)
// Worst time complexity:   O(n!)
// Space complexity:        O(1)
func Beeramid(bonus int, price float64) int {
	beers := int(float64(bonus) / price)

	for level := 1; ; level += 1 {
		square := level * level

		if beers < square {
			return level - 1
		}

		beers = beers - (square)
	}
}
