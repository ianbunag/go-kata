package product_of_array_except_self_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/3_product_of_array_except_self"
)

var _ = Describe("3ProductOfArrayExceptSelf", func() {
	It("should compute product of array except self", func() {
		Expect(ProductExceptSelf([]int{3, 2, 1})).To(Equal([]int{2, 3, 6}))
		Expect(ProductExceptSelf([]int{1, 2, 3, 4})).To(Equal([]int{24, 12, 8, 6}))
		Expect(ProductExceptSelf([]int{-1, 1, 0, -3, 3})).To(Equal([]int{0, 0, 9, 0, 0}))
	})
})
