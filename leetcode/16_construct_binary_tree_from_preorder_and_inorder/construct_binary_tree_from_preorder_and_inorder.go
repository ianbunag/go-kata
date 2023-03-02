package construct_binary_tree_from_preorder_and_inorder

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func BuildTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}

	rootInOrderIndex := 0

	for index, value := range inOrder {
		if value == preOrder[0] {
			rootInOrderIndex = index
			break
		}
	}

	return &TreeNode{
		Val:   preOrder[0],
		Left:  BuildTree(preOrder[1:rootInOrderIndex+1], inOrder[:rootInOrderIndex]),
		Right: BuildTree(preOrder[rootInOrderIndex+1:], inOrder[rootInOrderIndex+1:]),
	}
}
