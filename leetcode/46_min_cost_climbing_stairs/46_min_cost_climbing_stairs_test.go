package min_cost_climbing_stairs_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/46_min_cost_climbing_stairs"
)

var _ = Describe("46MinCostClimbingStairs", func() {
	It("it should return min cost of array with two values", func() {
		Expect(MinCostClimbingStairs([]int{10, 15})).To(Equal(10))
		Expect(MinCostClimbingStairs([]int{15, 10})).To(Equal(10))
	})

	It("it should return 15 on [10, 15, 20]", func() {
		Expect(MinCostClimbingStairs([]int{10, 15, 20})).To(Equal(15))
	})

	It("it should return 6 on [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]", func() {
		Expect(MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})).To(Equal(6))
	})
})
