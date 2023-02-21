package same_tree

import (
	. "github.com/yvnbunag/go-kata/lib/binary_tree"
)

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func IsSameTree(first *TreeNode, second *TreeNode) bool {
	firstQueue, secondQueue := []*TreeNode{first}, []*TreeNode{second}

	for len(firstQueue) != 0 {
		nextFirstQueue, nextSecondQueue := []*TreeNode{}, []*TreeNode{}

		for index, firstNode := range firstQueue {
			secondNode := secondQueue[index]

			if firstNode == secondNode {
				continue
			}

			if firstNode == nil || secondNode == nil || firstNode.Val != secondNode.Val {
				return false
			}

			nextFirstQueue = append(nextFirstQueue, firstNode.Left, firstNode.Right)
			nextSecondQueue = append(nextSecondQueue, secondNode.Left, secondNode.Right)
		}

		/**
		 * Assigning a new slice instead of editing in place consumes less memory as
		 *	slices contain underlying arrays under the hood.
		 */
		firstQueue, secondQueue = nextFirstQueue, nextSecondQueue
	}

	return true
}
