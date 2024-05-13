package max_area_of_island_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/69_max_area_of_island"
)

var _ = Describe("69MaxAreaOfIsland", func() {
	It("should return max area of island", func() {
		grid := [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}
		Expect(MaxAreaOfIsland(grid)).To(Equal(6))
	})

	It("should return 0 if there are no islands", func() {
		Expect(MaxAreaOfIsland([][]int{})).To(Equal(0))
		Expect(MaxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}})).To(Equal(0))
	})
})
