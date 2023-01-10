package longest_consecutive_sequence_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/4_longest_consecutive_sequence"
)

var _ = Describe("4LongestConsecutiveSequence", func() {
	It("should return longest consecutive sequence", func() {
		Expect(LongestConsecutive([]int{100, 4, 200, 1, 3, 2})).To(Equal(4))
		Expect(LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})).To(Equal(9))
		Expect(LongestConsecutive([]int{1, 3, 5})).To(Equal(1))
		Expect(LongestConsecutive([]int{})).To(Equal(0))
	})
})
