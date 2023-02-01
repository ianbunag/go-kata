package minimum_rotated_sorted_array_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/8_minimum_rotated_sorted_array"
)

var _ = Describe("8MinimumRotatedSortedArray", func() {
	It("should find minimum", func() {
		Expect(FindMin([]int{2, 3, 4, 5, 1})).To(Equal(1))
		Expect(FindMin([]int{2, 1})).To(Equal(1))
		Expect(FindMin([]int{1})).To(Equal(1))
		Expect(FindMin([]int{3, 4, 5, 1, 2})).To(Equal(1))
		Expect(FindMin([]int{4, 5, 6, 7, 0, 1, 2})).To(Equal(0))
		Expect(FindMin([]int{11, 13, 15, 17})).To(Equal(11))
	})
})
