package letter_combination_of_a_phone_number_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/64_letter_combination_of_a_phone_number"
)

var _ = Describe("64LetterCombinationOfAPhoneNumber", func() {
	It("should return empty array for empty string", func() {
		Expect(LetterCombinations("")).To(Equal([]string{}))
	})

	It("should return combination of a single digit string", func() {
		Expect(LetterCombinations("2")).To(Equal([]string{"a", "b", "c"}))
	})

	It("should return combination of a two digit string", func() {
		Expect(LetterCombinations("23")).To(Equal([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}))
	})
})
