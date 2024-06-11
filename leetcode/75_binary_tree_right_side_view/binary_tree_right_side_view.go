package binary_tree_right_side_view

const invalidValue = -101

// RightSideView
// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
func RightSideView(root *TreeNode) []int {
	result, queue := make([]int, 0), []*TreeNode{root}

	for layerLength := len(queue); layerLength > 0; layerLength = len(queue) {
		layerValue := invalidValue

		for _, node := range queue[0:layerLength] {
			if node != nil {
				layerValue, queue = node.Val, append(queue, node.Left, node.Right)
			}
		}

		result, queue = append(result, layerValue), queue[layerLength:]
	}

	if result[len(result)-1] == invalidValue {
		result = result[:len(result)-1]
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
