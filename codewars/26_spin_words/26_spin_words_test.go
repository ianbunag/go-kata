package spin_words_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/26_spin_words"
)

var _ = Describe("26SpinWords", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWords("Welcome")).To(Equal("emocleW"))
		Expect(SpinWords("to")).To(Equal("to"))
		Expect(SpinWords("CodeWars")).To(Equal("sraWedoC"))
	})
	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWords("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWords("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWords("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})
