package kth_largest_element_in_an_array_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/50_kth_largest_element_in_an_array"
)

var _ = Describe("50KthLargestElementInAnArray", func() {
	It("should return kth largest element in an array", func() {
		Expect(FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)).To(Equal(5))
		Expect(FindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)).To(Equal(4))
	})
})
