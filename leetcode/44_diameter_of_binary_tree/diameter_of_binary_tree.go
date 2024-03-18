package diameter_of_binary_tree

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func DiameterOfBinaryTree(root *TreeNode) int {
	rootDiameter, _ := extractLengths(root)

	return rootDiameter
}

func extractLengths(root *TreeNode) (int, int) {
	if root == nil {
		return -1, -1
	}

	leftDiameter, leftHeight := extractLengths(root.Left)
	rightDiameter, rightHeight := extractLengths(root.Right)

	return max(leftDiameter, rightDiameter, 2+leftHeight+rightHeight), 1 + max(leftHeight, rightHeight)
}
