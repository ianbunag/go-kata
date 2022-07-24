package reverse_words_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/1_reverse_words"
)

var _ = Describe("ReverseWord", func() {
	It("should reverse word", func() {
		Expect(ReverseWord("apple")).To(Equal("elppa"))
	})
})
