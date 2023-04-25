package maximum_product_subarray_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/32_maximum_product_subarray"
)

var _ = Describe("32MaximumProductSubarray", func() {
	It("should get maximum product subarray", func() {
		Expect(MaxProduct([]int{-2, 0, -1})).To(Equal(0))
		Expect(MaxProduct([]int{2, 3, -2, 4})).To(Equal(6))
		Expect(MaxProduct([]int{-4, -3})).To(Equal(12))
		Expect(MaxProduct([]int{2, 3, -2, 4, 0})).To(Equal(6))
	})
})
