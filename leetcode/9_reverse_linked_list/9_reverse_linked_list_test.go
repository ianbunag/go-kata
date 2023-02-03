package reverse_linked_list_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/9_reverse_linked_list"
)

var _ = Describe("9ReverseLinkedList", func() {
	It("should create new node from list", func() {
		head := NewNodeFromList([]int{1, 2, 3})

		Expect(head.Val).To(Equal(1))
		Expect(head.Next.Val).To(Equal(2))
		Expect(head.Next.Next.Val).To(Equal(3))
		Expect(head.Next.Next.Next).To(BeNil())
	})

	It("should convert node to list", func() {
		list := NewNodeFromList([]int{1, 2, 3})

		Expect(list.ToList()).To(Equal([]int{1, 2, 3}))
		Expect(list.Next.ToList()).To(Equal([]int{2, 3}))
	})

	It("should reverse linked list", func() {
		Expect(ReverseList(NewNodeFromList([]int{1, 2})).ToList()).To(Equal([]int{2, 1}))
		Expect(ReverseList(NewNodeFromList([]int{1, 2, 3, 4, 5})).ToList()).To(Equal([]int{5, 4, 3, 2, 1}))
		Expect(ReverseList(NewNodeFromList([]int{}))).To(BeNil())
	})
})
