package min_stack_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/52_min_stack"
)

var _ = Describe("52MinStack", func() {
	It("should get min stack", func() {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		Expect(minStack.GetMin()).To(Equal(-3))
		minStack.Pop()
		Expect(minStack.Top()).To(Equal(0))
		Expect(minStack.GetMin()).To(Equal(-2))
	})
})
