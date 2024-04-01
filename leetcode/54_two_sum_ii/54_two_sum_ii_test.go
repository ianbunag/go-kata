package two_sum_ii_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/54_two_sum_ii"
)

var _ = Describe("54TwoSumIi", func() {
	It("should get two sums", func() {
		Expect(TwoSum([]int{2, 7, 11, 15}, 9)).To(Equal([]int{1, 2}))
		Expect(TwoSum([]int{2, 7, 11, 15}, 22)).To(Equal([]int{2, 4}))
		Expect(TwoSum([]int{2, 3, 4}, 6)).To(Equal([]int{1, 3}))
		Expect(TwoSum([]int{-1, 0}, -1)).To(Equal([]int{1, 2}))
	})
})
