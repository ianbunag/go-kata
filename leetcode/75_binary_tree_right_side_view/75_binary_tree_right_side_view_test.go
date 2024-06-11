package binary_tree_right_side_view_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/75_binary_tree_right_side_view"
)

var _ = Describe("75BinaryTreeRightSideView", func() {
	It("should return binary tree right side view", func() {
		Expect(RightSideView(nil)).To(Equal([]int{}))

		nodes := []*TreeNode{
			{Val: 1},
			nil,
			{Val: 3},
		}
		nodes[0].Left = nodes[1]
		nodes[0].Right = nodes[2]
		Expect(RightSideView(nodes[0])).To(Equal([]int{1, 3}))

		nodes = []*TreeNode{
			{Val: 1},
			{Val: 2},
		}
		nodes[0].Left = nodes[1]
		Expect(RightSideView(nodes[0])).To(Equal([]int{1, 2}))

		nodes = []*TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			nil,
			{Val: 5},
			nil,
			{Val: 4},
		}
		nodes[0].Left = nodes[1]
		nodes[0].Right = nodes[2]
		nodes[1].Left = nodes[3]
		nodes[1].Right = nodes[4]
		nodes[2].Left = nodes[5]
		nodes[2].Right = nodes[6]
		Expect(RightSideView(nodes[0])).To(Equal([]int{1, 3, 4}))
	})
})
