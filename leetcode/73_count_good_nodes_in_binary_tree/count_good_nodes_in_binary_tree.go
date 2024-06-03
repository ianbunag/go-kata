package count_good_nodes_in_binary_tree

// GoodNodes
// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
func GoodNodes(root *TreeNode) int {
	return countGoodNodes(root, root.Val)
}

func countGoodNodes(node *TreeNode, maxValue int) int {
	if node == nil {
		return 0
	}

	if node.Val >= maxValue {
		return 1 + countGoodNodes(node.Left, node.Val) + countGoodNodes(node.Right, node.Val)
	}

	return countGoodNodes(node.Left, maxValue) + countGoodNodes(node.Right, maxValue)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
