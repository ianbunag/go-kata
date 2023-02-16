package invert_binary_tree

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	}

	return root
}
