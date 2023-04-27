package longest_increasing_subsequence_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/33_longest_increasing_subsequence"
)

var _ = Describe("33LongestIncreasingSubsequence", func() {
	It("should return length of longest increasing subsequence", func() {
		Expect(LengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})).To(Equal(6))
		Expect(LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})).To(Equal(4))
		Expect(LengthOfLIS([]int{0, 1, 0, 3, 2, 3})).To(Equal(4))
		Expect(LengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})).To(Equal(1))
	})
})
