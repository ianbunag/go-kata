package set_matrix_zeros_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/27_set_matrix_zeros"
)

var _ = Describe("27SetMatrixZeros", func() {
	It("should set matrix zeros", func() {
		simpleMatrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		simpleMatrixResult := [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}
		SetZeroes(simpleMatrix)
		Expect(simpleMatrix).To(Equal(simpleMatrixResult))

		regularMatrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		regularMatrixResult := [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}
		SetZeroes(regularMatrix)
		Expect(regularMatrix).To(Equal(regularMatrixResult))

		complexMatrix := [][]int{
			{1, 1, 1, 1, 0},
			{1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1},
		}
		complexMatrixResult := [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0},
		}
		SetZeroes(complexMatrix)
		Expect(complexMatrix).To(Equal(complexMatrixResult))
	})
})
