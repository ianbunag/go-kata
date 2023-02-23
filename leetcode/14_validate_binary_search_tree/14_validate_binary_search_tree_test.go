package validate_binary_search_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/14_validate_binary_search_tree"
	binary "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("14ValidateBinarySearchTree", func() {
	It("should return true for valid BST", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)
		node2.Left = node1
		node2.Right = node3

		Expect(IsValidBST(node2)).To(Equal(true))
	})

	It("should return false for invalid BST", func() {
		node1 := binary.New(1)
		node3 := binary.New(3)
		node4 := binary.New(4)
		node5 := binary.New(5)
		node6 := binary.New(6)
		node5.Left = node1
		node5.Right = node4
		node4.Left = node3
		node4.Right = node6

		Expect(IsValidBST(node5)).To(Equal(false))
	})

	It("should return false for invalid repeating value BST", func() {
		node1 := binary.New(2)
		node2 := binary.New(2)
		node3 := binary.New(2)
		node2.Left = node1
		node2.Right = node3

		Expect(IsValidBST(node2)).To(Equal(false))
	})

	It("should return false for imbalanced BST", func() {
		node3 := binary.New(3)
		node4 := binary.New(4)
		node5 := binary.New(5)
		node6 := binary.New(6)
		node7 := binary.New(7)
		node5.Left = node4
		node5.Right = node6
		node6.Left = node3
		node6.Right = node7

		Expect(IsValidBST(node5)).To(Equal(false))
	})
})
