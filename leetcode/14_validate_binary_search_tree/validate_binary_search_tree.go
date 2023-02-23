package validate_binary_search_tree

import (
	"math"

	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func IsValidBST(root *TreeNode) bool {
	return IsBalancedBST(root, math.MinInt64, math.MaxInt64)
}

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func IsBalancedBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return IsBalancedBST(root.Left, min, root.Val) && IsBalancedBST(root.Right, root.Val, max)
}
