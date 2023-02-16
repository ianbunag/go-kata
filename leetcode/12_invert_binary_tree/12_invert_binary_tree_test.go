package invert_binary_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/12_invert_binary_tree"
	binary "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("12InvertBinaryTree", func() {
	It("should invert simple binary tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)

		node2.Left = node1
		node2.Right = node3

		Expect(node2.ToList()).To(Equal([]int{1, 2, 3}))
		Expect(InvertTree(node2).ToList()).To(Equal([]int{3, 2, 1}))
	})

	It("should invert complex binary tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)
		node4 := binary.New(4)
		node6 := binary.New(6)
		node7 := binary.New(7)
		node9 := binary.New(9)

		node4.Left = node2
		node4.Right = node7
		node2.Left = node1
		node2.Right = node3
		node7.Left = node6
		node7.Right = node9

		Expect(node4.ToList()).To(Equal([]int{1, 2, 3, 4, 6, 7, 9}))
		Expect(InvertTree(node4).ToList()).To(Equal([]int{9, 7, 6, 4, 3, 2, 1}))
	})
})
