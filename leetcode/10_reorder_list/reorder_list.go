package reorder_list

import (
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func ReorderList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		next := head.Next
		head.Next = nil
		head = next
	}

	head = nodes[0]
	tail := head

	for index := 1; index < len(nodes); index += 1 {
		tailIndex := len(nodes) - index

		if index < tailIndex {
			tail.Next = nodes[tailIndex]
			tail = tail.Next
		}

		if index <= tailIndex {
			tail.Next = nodes[index]
			tail = tail.Next
		}
	}

	return head
}
