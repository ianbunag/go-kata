package binary_tree_max_path_sum_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/17_binary_tree_max_path_sum"
	binary "github.com/yvnbunag/go-kata/lib/binary_tree"
)

var _ = Describe("17BinaryTreeMaxPathSum", func() {
	It("should return max path sum from a simple tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node3 := binary.New(3)
		node1.Left = node2
		node1.Right = node3

		Expect(MaxPathSum(node1)).To(Equal(6))
	})

	It("should return max path sum from a standard tree", func() {
		nodeN10 := binary.New(-10)
		node9 := binary.New(9)
		node20 := binary.New(20)
		node15 := binary.New(15)
		node7 := binary.New(7)
		nodeN10.Left = node9
		nodeN10.Right = node20
		node20.Left = node15
		node20.Right = node7

		Expect(MaxPathSum(nodeN10)).To(Equal(42))
	})

	It("should return max path sum from the root of a straight tree", func() {
		node2 := binary.New(2)
		nodeN1 := binary.New(-1)
		node2.Left = nodeN1

		Expect(MaxPathSum(node2)).To(Equal(2))
	})

	It("should return max path sum from a straight tree", func() {
		node1 := binary.New(1)
		node2 := binary.New(2)
		node1.Left = node2

		Expect(MaxPathSum(node1)).To(Equal(3))
	})

	It("should return max path sum from a single node", func() {
		nodeN3 := binary.New(-3)

		Expect(MaxPathSum(nodeN3)).To(Equal(-3))
	})

	It("should return max path sum from a straight node of a simple tree", func() {
		node1 := binary.New(1)
		nodeN2 := binary.New(-2)
		node3 := binary.New(3)
		node1.Left = nodeN2
		node1.Right = node3

		Expect(MaxPathSum(node1)).To(Equal(4))
	})

	It("should return max path sum from a single leaf node of a straight tree", func() {
		nodeN2 := binary.New(-2)
		nodeN1 := binary.New(-1)
		nodeN2.Left = nodeN1

		Expect(MaxPathSum(nodeN2)).To(Equal(-1))
	})

	It("should return max path sum from a single leaf node of a standard tree", func() {
		node := binary.New(1)
		node.Left = binary.New(-2)
		node.Right = binary.New(-3)
		node.Left.Left = binary.New(1)
		node.Left.Right = binary.New(3)
		node.Right.Left = binary.New(-2)
		node.Left.Left.Left = binary.New(-1)

		Expect(MaxPathSum(node)).To(Equal(3))
	})
})
