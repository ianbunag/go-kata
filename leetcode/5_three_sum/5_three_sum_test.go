package three_sum_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/5_three_sum"
)

var _ = Describe("5ThreeSum", func() {
	It("should return three sum", func() {
		Expect(GetThreeSum([]int{})).To(Equal([][]int{}))
		Expect(GetThreeSum([]int{0, 1, 2, 3})).To(Equal([][]int{}))
		Expect(GetThreeSum([]int{-1, 0, 1, 2, -1, -4})).To(Equal([][]int{{-1, -1, 2}, {-1, 0, 1}}))
		Expect(GetThreeSum([]int{1, -1, -1, 0})).To(Equal([][]int{{-1, 0, 1}}))
		Expect(GetThreeSum([]int{0, 1, 1})).To(Equal([][]int{}))
		Expect(GetThreeSum([]int{0, 0, 0})).To(Equal([][]int{{0, 0, 0}}))
		Expect(GetThreeSum(make([]int, 5000))).To(Equal([][]int{{0, 0, 0}}))
	})
})
