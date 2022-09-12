package maximum_subarray_sum_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/29_maximum_subarray_sum"
)

var _ = Describe("29MaximumSubarraySum", func() {
	It("Basic tests", func() {
		Expect(MaximumSubarraySum([]int{})).To(Equal(0))
		Expect(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})).To(Equal(6))
		Expect(MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})).To(Equal(0))
	})
})
