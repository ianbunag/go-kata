package palindromic_substrings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/30_palindromic_substrings"
)

var _ = Describe("30PalindromicSubstrings", func() {
	It("should count substrings", func() {
		Expect(CountSubstrings("abc")).To(Equal(3))
		Expect(CountSubstrings("aaa")).To(Equal(6))
	})
})
