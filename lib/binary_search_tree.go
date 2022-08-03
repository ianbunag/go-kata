package lib

import (
	"errors"
)

type TreeNode struct {
	Value  int
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

func (treeNode *TreeNode) Insert(value int) error {
	if treeNode == nil {
		return errors.New("Tree is nil")
	}

	if treeNode.Value == value {
		return errors.New("Node exists")
	}

	if treeNode.Value > value {
		if treeNode.left == nil {
			treeNode.left = &TreeNode{parent: treeNode, Value: value}
			return nil
		}

		return treeNode.left.Insert(value)
	}

	if treeNode.Value < value {
		if treeNode.right == nil {
			treeNode.right = &TreeNode{parent: treeNode, Value: value}
			return nil
		}
		return treeNode.right.Insert(value)
	}

	return nil
}

func (treeNode *TreeNode) ToList() (result []int) {
	if treeNode.left != nil {
		result = append(result, treeNode.left.ToList()...)
	}

	result = append(result, treeNode.Value)

	if treeNode.right != nil {
		result = append(result, treeNode.right.ToList()...)
	}

	return
}

func (treeNode *TreeNode) Get(value int) *TreeNode {
	if treeNode.Value == value {
		return treeNode
	}

	if treeNode.Value > value && treeNode.left != nil {
		return treeNode.left.Get(value)
	}

	if treeNode.Value < value && treeNode.right != nil {
		return treeNode.right.Get(value)
	}

	return nil
}

func (treeNode *TreeNode) Prev() *TreeNode {
	if treeNode.left != nil {
		return treeNode.left
	}

	if treeNode.parent.Value < treeNode.Value {
		return treeNode.parent
	}

	return nil
}

func (treeNode *TreeNode) Next() *TreeNode {
	if treeNode.right != nil {
		return treeNode.right
	}

	if treeNode.parent.Value > treeNode.Value {
		return treeNode.parent
	}

	return nil
}
