package binary_tree_max_path_sum

import (
	"math"

	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func MaxPathSum(root *TreeNode) int {
	return getNodeSums(root).maxPath
}

func getNodeSums(root *TreeNode) NodeSums {
	current := NodeSums{
		linearMaxPath: math.MinInt32,
		maxPath:       math.MinInt32,
	}

	if root == nil {
		return current
	}

	left, right := getNodeSums(root.Left), getNodeSums(root.Right)
	current.linearMaxPath = max(
		root.Val,
		root.Val+left.linearMaxPath,
		root.Val+right.linearMaxPath,
	)
	current.maxPath = max(
		current.linearMaxPath,
		root.Val+left.linearMaxPath+right.linearMaxPath,
		left.maxPath,
		right.maxPath,
	)

	return current
}

type NodeSums struct {
	linearMaxPath int
	maxPath       int
}

func max(maxValue int, values ...int) int {
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}
