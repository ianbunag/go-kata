package group_anagrams_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/2_group_anagrams"
)

var _ = Describe("2GroupAnagrams", func() {
	It("should group single string", func() {
		result := GroupAnagrams([]string{"a"})
		expected := [][]string{
			{"a"},
		}
		Expect(result).To(ConsistOf(expected))
	})

	It("should group empty string", func() {
		result := GroupAnagrams([]string{""})
		expected := [][]string{
			{""},
		}
		Expect(result).To(ConsistOf(expected))
	})

	It("should group anagrams", func() {
		result := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
		expected := [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}
		Expect(result).To(ConsistOf(expected))
	})
})
