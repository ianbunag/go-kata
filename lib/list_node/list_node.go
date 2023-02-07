package lib_list_node

func NewNodeFromList(values []int) *ListNode {
	if len(values) < 1 {
		return nil
	}

	head := &ListNode{values[0], nil}
	currentNode := head

	for _, value := range values[1:] {
		currentNode.Next = &ListNode{value, nil}
		currentNode = currentNode.Next
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNode *ListNode) ToList() []int {
	list := []int{}
	currentNode := listNode

	for {
		list = append(list, currentNode.Val)

		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}

	return list
}
