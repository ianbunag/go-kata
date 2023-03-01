package kth_smallest_element_in_bst_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/15_kth_smallest_element_in_bst"
	binary "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("15KthSmallestElementInBst", func() {
	It("should get smallest element from simple tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)
		node4 := binary.New(4)
		node3.Left = node1
		node3.Right = node4
		node1.Right = node2

		Expect(KthSmallest(node3, 1)).To(Equal(1))
		Expect(KthSmallest(node3, 2)).To(Equal(2))
		Expect(KthSmallest(node3, 3)).To(Equal(3))
		Expect(KthSmallest(node3, 4)).To(Equal(4))
	})

	It("should get smallest element from complex tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)
		node4 := binary.New(4)
		node5 := binary.New(5)
		node6 := binary.New(6)
		node5.Left = node3
		node5.Right = node6
		node3.Left = node2
		node3.Right = node4
		node2.Left = node1

		Expect(KthSmallest(node5, 1)).To(Equal(1))
		Expect(KthSmallest(node5, 2)).To(Equal(2))
		Expect(KthSmallest(node5, 3)).To(Equal(3))
		Expect(KthSmallest(node5, 4)).To(Equal(4))
		Expect(KthSmallest(node5, 5)).To(Equal(5))
		Expect(KthSmallest(node5, 6)).To(Equal(6))
	})
})
