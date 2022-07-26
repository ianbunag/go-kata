package array_diff_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/5_array_diff"
)

var _ = Describe("ArrayDiff", func() {
	It("should handle basic cases", func() {
		var emptyArr []int
		Expect(ArrayDiff(emptyArr, []int{1, 2})).To(BeEmpty())
		Expect(ArrayDiff([]int{1, 2, 2}, emptyArr)).To(Equal([]int{1, 2, 2}))
		Expect(ArrayDiff([]int{1, 2}, []int{1})).To(Equal([]int{2}))
		Expect(ArrayDiff([]int{1, 2, 2}, []int{1})).To(Equal([]int{2, 2}))
		Expect(ArrayDiff([]int{1, 2, 2}, []int{2})).To(Equal([]int{1}))
		Expect(ArrayDiff([]int{1, 2, 3}, []int{1, 2})).To(Equal([]int{3}))
	})
})
