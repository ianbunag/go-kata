package anagrams_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/23_anagrams"
)

func dotest(word string, words, expected []string) {
	actual := Anagrams(word, words)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty())
	} else {
		Expect(actual).To(Equal(expected))
	}
}

var _ = Describe("23Anagrams", func() {
	It("should handle basic cases", func() {
		dotest("laser", []string{"lazing", "lazy", "lacer"}, nil)
		dotest("abba", []string{"aabb", "abcd", "bbaa", "dada"}, []string{"aabb", "bbaa"})
		dotest("racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"}, []string{"carer", "arcre", "carre"})
	})
})
