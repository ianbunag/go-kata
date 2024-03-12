package search_a_2d_matrix_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/43_search_a_2d_matrix"
)

var _ = Describe("43SearchA2dMatrix", func() {
	It("should search target within an empty matrix", func() {
		Expect(SearchMatrix([][]int{}, 5)).To(BeFalse())
	})

	It("should search target within a matrix with a single row", func() {
		Expect(SearchMatrix([][]int{{5}}, 5)).To(BeTrue())
		Expect(SearchMatrix([][]int{{5}}, -5)).To(BeFalse())
	})

	It("should return true if target is found in the matrix", func() {
		Expect(SearchMatrix([][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 3)).To(BeTrue())
	})

	It("should return false if target is not found in the matrix", func() {
		Expect(SearchMatrix([][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 13)).To(BeFalse())
	})
})
