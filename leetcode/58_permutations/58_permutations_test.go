package permutations_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/58_permutations"
)

var _ = Describe("58Permutations", func() {
	It("should return all possible permutations", func() {
		Expect(Permute([]int{1})).To(ConsistOf([][]int{
			{1},
		}))
		Expect(Permute([]int{0, 1})).To(ConsistOf([][]int{
			{0, 1},
			{1, 0},
		}))
		Expect(Permute([]int{1, 2, 3})).To(ConsistOf([][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}))
	})
})
