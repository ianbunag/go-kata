package interleaving_string_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/72_interleaving_string"
)

var _ = Describe("72InterleavingString", func() {
	It("should return true if s3 is formed by interleaving s1 and s2", func() {
		Expect(IsInterleave("aabcc", "dbbca", "aadbbcbcac")).To(Equal(true))
		Expect(IsInterleave("", "", "")).To(Equal(true))
		Expect(IsInterleave("", "a", "a")).To(Equal(true))
		Expect(IsInterleave("a", "", "a")).To(Equal(true))
	})

	It("should return false if s3 cannot be formed by interleaving s1 and s2", func() {
		Expect(IsInterleave("", "", "a")).To(Equal(false))
		Expect(IsInterleave("", "a", "")).To(Equal(false))
		Expect(IsInterleave("a", "", "")).To(Equal(false))
		Expect(IsInterleave("a", "a", "")).To(Equal(false))
		Expect(IsInterleave("aabcc", "dbbca", "aadbbbaccc")).To(Equal(false))
		Expect(IsInterleave(
			"aaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		)).To(Equal(false))
		Expect(IsInterleave(
			"abababababababababababababababababababababababababababababababababababababababababababababababababbb",
			"babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
			"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
		)).To(Equal(false))
	})
})
