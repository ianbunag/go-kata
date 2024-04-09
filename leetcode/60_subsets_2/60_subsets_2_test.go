package subsets_2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/60_subsets_2"
)

var _ = Describe("60Subsets2", func() {
	It("should get subsets with duplicates", func() {
		Expect(SubsetsWithDup([]int{1})).To(ConsistOf([][]int{
			{},
			{1},
		}))
		Expect(SubsetsWithDup([]int{1, 2, 3})).To(ConsistOf([][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}))
		Expect(SubsetsWithDup([]int{1, 2, 2})).To(ConsistOf([][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 2},
			{2},
			{2, 2},
		}))
		Expect(SubsetsWithDup([]int{4, 4, 4, 1, 4})).To(ConsistOf([][]int{
			{},
			{1},
			{1, 4},
			{1, 4, 4},
			{1, 4, 4, 4},
			{1, 4, 4, 4, 4},
			{4},
			{4, 4},
			{4, 4, 4},
			{4, 4, 4, 4},
		}))
		Expect(SubsetsWithDup([]int{2, 1, 2, 1, 3})).To(ConsistOf([][]int{
			{},
			{1},
			{1, 1},
			{1, 1, 2},
			{1, 1, 2, 2},
			{1, 1, 2, 2, 3},
			{1, 1, 2, 3},
			{1, 1, 3},
			{1, 2},
			{1, 2, 2},
			{1, 2, 2, 3},
			{1, 2, 3},
			{1, 3},
			{2},
			{2, 2},
			{2, 2, 3},
			{2, 3},
			{3},
		}))
	})
})
