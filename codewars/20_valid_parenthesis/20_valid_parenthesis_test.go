package valid_parenthesis_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/20_valid_parenthesis"
)

var _ = Describe("20ValidParenthesis", func() {
	It("passes example tests", func() {
		Expect(IsValidParentheses("()")).To(Equal(true))
		Expect(IsValidParentheses(")")).To(Equal(false))
		Expect(IsValidParentheses(")(()))")).To(Equal(false))
		Expect(IsValidParentheses("(")).To(Equal(false))
		Expect(IsValidParentheses("(())((()())())")).To(Equal(true))
	})
})
