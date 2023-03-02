package construct_binary_tree_from_preorder_and_inorder_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/leetcode/16_construct_binary_tree_from_preorder_and_inorder"
)

var _ = Describe("16ConstructBinaryTreeFromPreorderAndInorder", func() {
	It("should contruct binary tree from preorder and inorder traversal list", func() {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		tree := BuildTree(preorder, inorder)

		Expect(tree.Val).To(Equal(3))
		Expect(tree.Left.Val).To(Equal(9))
		Expect(tree.Right.Val).To(Equal(20))
		Expect(tree.Right.Left.Val).To(Equal(15))
		Expect(tree.Right.Right.Val).To(Equal(7))
	})
})
