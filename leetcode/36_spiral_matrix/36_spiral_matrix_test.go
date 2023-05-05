package spiral_matrix_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/36_spiral_matrix"
)

var _ = Describe("36SpiralMatrix", func() {
	It("should return spiral order", func() {
		Expect(SpiralOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})).To(Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}))
		Expect(SpiralOrder([][]int{
			{2, 3},
		})).To(Equal([]int{2, 3}))
		Expect(SpiralOrder([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		})).To(Equal([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}))
	})
})
