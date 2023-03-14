package implement_trie_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/18_implement_trie"
)

var _ = Describe("18ImplementTrie", func() {
	It("should store and retrieve values", func() {
		trie := Constructor()
		trie.Insert("apple")
		Expect(trie.Search("apple")).To(Equal(true))
		Expect(trie.Search("app")).To(Equal(false))
		Expect(trie.StartsWith("app")).To(Equal(true))
		trie.Insert("app")
		Expect(trie.Search("app")).To(Equal(true))
	})
})
