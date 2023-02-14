package linked_list_cycle_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/11_linked_list_cycle"
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

var _ = Describe("11LinkedListCycle", func() {
	It("should determine if linked list has cycle", func() {
		flatHead := &ListNode{Val: 1, Next: nil}

		Expect(HasCycle(flatHead)).To(Equal(false))

		simpleHead := &ListNode{Val: 1, Next: nil}
		simpleHead.Next = &ListNode{Val: 2, Next: nil}
		simpleHead.Next.Next = simpleHead

		Expect(HasCycle(simpleHead)).To(Equal(true))

		complexHead := &ListNode{Val: 3, Next: nil}
		complexHead.Next = &ListNode{Val: 2, Next: nil}
		complexHead.Next.Next = &ListNode{Val: 0, Next: nil}
		complexHead.Next.Next.Next = &ListNode{Val: -4, Next: nil}
		complexHead.Next.Next.Next.Next = complexHead.Next

		Expect(HasCycle(complexHead)).To(Equal(true))
	})
})
