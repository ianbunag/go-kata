package count_good_nodes_in_binary_tree_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/73_count_good_nodes_in_binary_tree"
)

var _ = Describe("73CountGoodNodesInBinaryTree", func() {
	It("should return number of good nodes", func() {
		nodes := []*TreeNode{
			{Val: 3},
			{Val: 1},
			{Val: 4},
			{Val: 3},
			nil,
			{Val: 1},
			{Val: 5},
		}
		nodes[0].Left = nodes[1]
		nodes[0].Right = nodes[2]
		nodes[1].Left = nodes[3]
		nodes[1].Right = nodes[4]
		nodes[2].Left = nodes[5]
		nodes[2].Right = nodes[6]
		Expect(GoodNodes(nodes[0])).To(Equal(4))

		nodes = []*TreeNode{
			{Val: 3},
			{Val: 3},
			nil,
			{Val: 4},
			{Val: 2},
		}
		nodes[0].Left = nodes[1]
		nodes[0].Right = nodes[2]
		nodes[1].Left = nodes[3]
		nodes[1].Right = nodes[4]
		Expect(GoodNodes(nodes[0])).To(Equal(3))

		nodes = []*TreeNode{
			{Val: 1},
		}
		Expect(GoodNodes(nodes[0])).To(Equal(1))
	})
})
