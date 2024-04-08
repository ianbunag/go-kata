package add_two_numbers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/56_add_two_numbers"
	ListNode "github.com/yvnbunag/go-kata/lib/list_node"
)

var _ = Describe("56AddTwoNumbers", func() {
	It("should add two numbers", func() {
		l1 := ListNode.NewNodeFromList([]int{2, 4, 3})
		l2 := ListNode.NewNodeFromList([]int{5, 6, 4})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{7, 0, 8}))

		l1 = ListNode.NewNodeFromList([]int{0})
		l2 = ListNode.NewNodeFromList([]int{0})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{0}))

		l1 = ListNode.NewNodeFromList([]int{9})
		l2 = ListNode.NewNodeFromList([]int{9})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{8, 1}))

		l1 = ListNode.NewNodeFromList([]int{9, 9, 9, 9, 9, 9, 9})
		l2 = ListNode.NewNodeFromList([]int{9, 9, 9, 9})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{8, 9, 9, 9, 0, 0, 0, 1}))

		l1 = ListNode.NewNodeFromList([]int{0, 0, 1})
		l2 = ListNode.NewNodeFromList([]int{1})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{1, 0, 1}))

		l1 = ListNode.NewNodeFromList([]int{1})
		l2 = ListNode.NewNodeFromList([]int{0, 0, 1})
		Expect(AddTwoNumbers(l1, l2).ToList()).To(Equal([]int{1, 0, 1}))
	})
})
