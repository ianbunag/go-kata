package alien_dictionary_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/28_alien_dictionary"
)

var _ = Describe("28AlienDictionary", func() {
	It("should get order", func() {
		Expect(
			AlienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}),
		).To(Equal("wertf"))
		Expect(AlienOrder([]string{"z", "x"})).To(Equal("zx"))
		Expect(AlienOrder([]string{"a", "ac", "bc", "c"})).To(Equal("abc"))
	})

	It("should return empty string on invalid orders", func() {
		Expect(AlienOrder([]string{"ab", "a"})).To(Equal(""))
	})
})
