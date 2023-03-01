package kth_smallest_element_in_bst

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(log n)
func KthSmallest(root *TreeNode, k int) int {
	var smallest *TreeNode
	stack := extractLeftNodes(root)

	for ; k > 0; k -= 1 {
		stackLastIndex := len(stack) - 1
		smallest = stack[stackLastIndex]
		stack = append(stack[:stackLastIndex], extractLeftNodes(smallest.Right)...)
	}

	return smallest.Val
}

func extractLeftNodes(root *TreeNode) []*TreeNode {
	leftNodes := []*TreeNode{}

	for root != nil {
		leftNodes = append(leftNodes, root)
		root = root.Left
	}

	return leftNodes
}
