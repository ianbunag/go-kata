package lib_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(Val int) *TreeNode {
	return &TreeNode{Val: Val}
}

func (treeNode *TreeNode) ToList() (result []int) {
	if treeNode.Left != nil {
		result = append(result, treeNode.Left.ToList()...)
	}

	result = append(result, treeNode.Val)

	if treeNode.Right != nil {
		result = append(result, treeNode.Right.ToList()...)
	}

	return
}
