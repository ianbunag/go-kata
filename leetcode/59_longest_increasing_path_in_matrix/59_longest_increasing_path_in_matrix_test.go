package longest_increasing_path_in_matrix_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/59_longest_increasing_path_in_matrix"
)

var _ = Describe("59LongestIncreasingPathInMatrix", func() {
	It("should return longest increasing path", func() {
		Expect(LongestIncreasingPath([][]int{{1}})).To(Equal(1))
		Expect(LongestIncreasingPath([][]int{{1, 2}})).To(Equal(2))
		Expect(LongestIncreasingPath([][]int{
			{1, 3},
			{3, 2},
		})).To(Equal(2))
		Expect(LongestIncreasingPath([][]int{
			{9, 9, 4},
			{6, 6, 8},
			{2, 1, 1},
		})).To(Equal(4))
		Expect(LongestIncreasingPath([][]int{
			{3, 4, 5},
			{3, 2, 6},
			{2, 2, 1},
		})).To(Equal(4))
	})
})
