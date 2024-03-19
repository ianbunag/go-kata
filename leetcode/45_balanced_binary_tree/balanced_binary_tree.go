package balanced_binary_tree

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func IsBalanced(root *TreeNode) bool {
	balanced, _ := getBalanceAndHeight(root)

	return balanced
}

func getBalanceAndHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftHeight := getBalanceAndHeight(root.Left)

	if !leftBalanced {
		return false, 0
	}

	rightBalanced, rightHeight := getBalanceAndHeight(root.Right)

	if !rightBalanced || abs(leftHeight-rightHeight) > 1 {
		return false, 0
	}

	return true, 1 + max(leftHeight, rightHeight)
}

func abs(value int) int {
	if value > 0 {
		return value
	}

	return -value
}
