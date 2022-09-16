package first_non_repeating_character_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/33_first_non_repeating_character"
)

var _ = Describe("33FirstNonRepeatingCharacter", func() {
	It("should handle simple tests", func() {
		Expect(FirstNonRepeating("a")).To(Equal("a"))
		Expect(FirstNonRepeating("stress")).To(Equal("t"))
		Expect(FirstNonRepeating("moonmen")).To(Equal("e"))
	})
	It("should handle empty strings", func() {
		Expect(FirstNonRepeating("")).To(Equal(""))
	})
	It("should handle all repeating strings", func() {
		Expect(FirstNonRepeating("abba")).To(Equal(""))
		Expect(FirstNonRepeating("aa")).To(Equal(""))
	})
	It("should handle odd characters", func() {
		Expect(FirstNonRepeating("~><#~><")).To(Equal("#"))
		Expect(FirstNonRepeating("hello world, eh?")).To(Equal("w"))
	})
	It("should handle letter cases", func() {
		Expect(FirstNonRepeating("sTreSS")).To(Equal("T"))
		Expect(FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")).To(Equal(","))
	})
})
