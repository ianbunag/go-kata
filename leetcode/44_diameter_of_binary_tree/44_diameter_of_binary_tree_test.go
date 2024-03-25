package diameter_of_binary_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/44_diameter_of_binary_tree"
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("44DiameterOfBinaryTree", func() {
	It("should return -1 on nil tree", func() {
		Expect(DiameterOfBinaryTree(nil)).To(Equal(-1))
	})

	It("should return 0 on single node tree", func() {
		node1 := New(1)

		Expect(DiameterOfBinaryTree(node1)).To(Equal(0))
	})

	It("should return diameter of a simple tree", func() {
		node1 := New(1)
		node2 := New(2)

		node1.Left = node2

		Expect(DiameterOfBinaryTree(node1)).To(Equal(1))
	})

	It("should return diameter of a complex tree", func() {
		node1 := New(1)
		node2 := New(2)
		node3 := New(3)
		node4 := New(4)
		node5 := New(5)

		node1.Left = node2
		node1.Right = node3
		node2.Left = node4
		node2.Right = node5

		Expect(DiameterOfBinaryTree(node1)).To(Equal(3))
	})
})
