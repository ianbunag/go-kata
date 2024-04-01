package evaluate_reverse_polish_notation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/53_evaluate_reverse_polish_notation"
)

var _ = Describe("53EvaluateReversePolishNotation", func() {
	It("should evaluate reverse polish notation", func() {
		Expect(EvalRPN([]string{"2", "1", "+"})).To(Equal(3))
		Expect(EvalRPN([]string{"2", "1", "+", "3", "*"})).To(Equal(9))
		Expect(EvalRPN([]string{"4", "13", "5", "/", "+"})).To(Equal(6))
		Expect(EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})).To(Equal(22))
	})
})
