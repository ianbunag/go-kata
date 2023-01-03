package valid_anagram_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/1_valid_anagram"
)

var _ = Describe("1ValidAnagram", func() {
	It("should check if words are an anagram of each other", func() {
		Expect(IsAnagram("anagram", "nagaram")).To(Equal(true))
		Expect(IsAnagram("rat", "car")).To(Equal(false))
	})
})
