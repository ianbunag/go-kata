package minimum_window_substring_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/7_minimum_window_substring"
)

var _ = Describe("7MinimumWindowSubstring", func() {
	It("should get minimum window substring", func() {
		Expect(MinWindow("cabwefgewcwaefgcf", "cae")).To(Equal("cwae"))
		Expect(MinWindow("ADOBECODECBANC", "ABCC")).To(Equal("CBANC"))
		Expect(MinWindow("ADOBECODEBANC", "ABC")).To(Equal("BANC"))
		Expect(MinWindow("a", "a")).To(Equal("a"))
		Expect(MinWindow("a", "aa")).To(Equal(""))
	})
})
