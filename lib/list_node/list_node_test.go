package lib_list_node_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

var _ = Describe("ListNode", func() {
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
})
