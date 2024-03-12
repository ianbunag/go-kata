package binary_search_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/42_binary_search"
)

var _ = Describe("42BinarySearch", func() {
	It("should search target within an empty array", func() {
		Expect(Search([]int{}, 5)).To(Equal(-1))
	})

	It("should search target within an array with a single element", func() {
		Expect(Search([]int{5}, 5)).To(Equal(0))
		Expect(Search([]int{5}, -5)).To(Equal(-1))
	})

	It("should search target within an array with an odd number of elements", func() {
		Expect(Search([]int{1, 2, 3}, 1)).To(Equal(0))
		Expect(Search([]int{1, 2, 3}, 2)).To(Equal(1))
		Expect(Search([]int{1, 2, 3}, 3)).To(Equal(2))
		Expect(Search([]int{1, 2, 3}, 0)).To(Equal(-1))
		Expect(Search([]int{1, 2, 3}, 4)).To(Equal(-1))
	})

	It("should search target within an array with an even number of elements", func() {
		Expect(Search([]int{1, 2, 3, 4}, 1)).To(Equal(0))
		Expect(Search([]int{1, 2, 3, 4}, 2)).To(Equal(1))
		Expect(Search([]int{1, 2, 3, 4}, 3)).To(Equal(2))
		Expect(Search([]int{1, 2, 3, 4}, 4)).To(Equal(3))
		Expect(Search([]int{1, 2, 3, 4}, 0)).To(Equal(-1))
		Expect(Search([]int{1, 2, 3, 4}, 5)).To(Equal(-1))
	})
})
