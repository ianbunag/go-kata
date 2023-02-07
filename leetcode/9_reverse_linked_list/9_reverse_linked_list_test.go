package reverse_linked_list_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/9_reverse_linked_list"
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

var _ = Describe("9ReverseLinkedList", func() {
	It("should reverse linked list", func() {
		Expect(ReverseList(NewNodeFromList([]int{1, 2})).ToList()).To(Equal([]int{2, 1}))
		Expect(ReverseList(NewNodeFromList([]int{1, 2, 3, 4, 5})).ToList()).To(Equal([]int{5, 4, 3, 2, 1}))
		Expect(ReverseList(NewNodeFromList([]int{}))).To(BeNil())
	})
})
