package permutations_test

import (
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/39_permutations"
)

func doTest(s string, expected []string) {
	actual := Permutations(s)
	sort.Strings(actual)
	Expect(actual).To(Equal(expected))
}

var _ = Describe("39Permutations", func() {
	It("Sample tests - Unique letters", func() {
		// doTest("a", []string{"a"})
		doTest("ab", []string{"ab", "ba"})
		doTest("abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"})
		abcd := []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
			"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
			"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
			"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"}
		doTest("abcd", abcd)
		doTest("bcda", abcd)
		doTest("dcba", abcd)
	})
	It("Sample tests - Duplicate letters", func() {
		doTest("aa", []string{"aa"})
		doTest("aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"})
		doTest("aaaab", []string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"})
		doTest("aaaaab", []string{"aaaaab", "aaaaba", "aaabaa", "aabaaa", "abaaaa", "baaaaa"})
	})
})
