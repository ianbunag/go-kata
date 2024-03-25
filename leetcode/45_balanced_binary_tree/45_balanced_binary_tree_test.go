package balanced_binary_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/45_balanced_binary_tree"
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("45BalancedBinaryTree", func() {
	It("should return true on nil tree", func() {
		Expect(IsBalanced(nil)).To(Equal(true))
	})

	It("should return true on single node tree", func() {
		node1 := New(1)

		Expect(IsBalanced(node1)).To(Equal(true))
	})

	It("should return true on balanced tree", func() {
		node3 := New(3)
		node7 := New(7)
		node9 := New(9)
		node15 := New(15)
		node20 := New(20)

		node3.Left = node9
		node3.Right = node20
		node20.Left = node15
		node20.Right = node7

		Expect(IsBalanced(node3)).To(Equal(true))
	})

	It("should return false on unbalanced tree", func() {
		node1 := New(1)
		node2A := New(2)
		node2B := New(2)
		node3A := New(3)
		node3B := New(3)
		node4A := New(4)
		node4B := New(4)

		node1.Left = node2A
		node1.Right = node2B
		node2A.Left = node3A
		node2A.Right = node3B
		node3A.Left = node4A
		node3A.Right = node4B

		Expect(IsBalanced(node1)).To(Equal(false))
	})
})
