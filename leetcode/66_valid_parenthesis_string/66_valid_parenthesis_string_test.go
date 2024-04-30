package valid_parenthesis_string_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/66_valid_parenthesis_string"
)

var _ = Describe("66ValidParenthesisString", func() {
	It("should return true for valid parenthesis string", func() {
		Expect(CheckValidString("()")).To(Equal(true))
		Expect(CheckValidString("()()")).To(Equal(true))
		Expect(CheckValidString("(()())")).To(Equal(true))
		Expect(CheckValidString("(*)")).To(Equal(true))
		Expect(CheckValidString("(*))")).To(Equal(true))
		Expect(CheckValidString("************************************************************")).To(Equal(true))
		Expect(CheckValidString("**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))")).To(Equal(true))
	})

	It("should return false for invalid parenthesis string", func() {
		Expect(CheckValidString("(*()))*(")).To(Equal(false))
		Expect(CheckValidString("(")).To(Equal(false))
		Expect(CheckValidString(")")).To(Equal(false))
		Expect(CheckValidString(")(")).To(Equal(false))
		Expect(CheckValidString("())(")).To(Equal(false))
	})
})
