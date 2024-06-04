package permutation_string_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/74_permutation_string"
)

var _ = Describe("74PermutationString", func() {
	It("should return true if s2 contains a permutation of s1", func() {
		Expect(CheckInclusion("ab", "eidbaooo")).To(Equal(true))
		Expect(CheckInclusion("a", "ab")).To(Equal(true))
		Expect(CheckInclusion("adc", "dcda")).To(Equal(true))
		Expect(CheckInclusion("abc", "bbbca")).To(Equal(true))
		Expect(CheckInclusion("ab", "a")).To(Equal(false))
		Expect(CheckInclusion("ab", "eidboaoo")).To(Equal(false))
	})
})
