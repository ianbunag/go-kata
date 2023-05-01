package longest_common_subsequence_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/34_longest_common_subsequence"
)

var _ = Describe("34LongestCommonSubsequence", func() {
	It("should get longest common subsequence", func() {
		Expect(LongestCommonSubsequence("abcde", "ace")).To(Equal(3))
		Expect(LongestCommonSubsequence("bsbininm", "jmjkbkjkv")).To(Equal(1))
		Expect(LongestCommonSubsequence("", "")).To(Equal(0))
		Expect(LongestCommonSubsequence("a", "a")).To(Equal(1))
		Expect(LongestCommonSubsequence("abcde", "acce")).To(Equal(3))
		Expect(LongestCommonSubsequence("abcde", "afcge")).To(Equal(3))
		Expect(LongestCommonSubsequence("abccde", "afccge")).To(Equal(4))
		Expect(LongestCommonSubsequence("abc", "abc")).To(Equal(3))
		Expect(LongestCommonSubsequence("abc", "def")).To(Equal(0))
	})
})
