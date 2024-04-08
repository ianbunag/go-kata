package add_two_numbers

import (
	listnode "github.com/yvnbunag/go-kata/lib/list_node"
)

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func AddTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	head, tail, carry := l1, l1, 0

	for l1 != nil && l2 != nil {
		carry += l1.Val + l2.Val
		carry, l1.Val, tail, l1, l2 = carry/10, carry%10, l1, l1.Next, l2.Next
	}

	if l2 != nil {
		tail.Next = l2
	}

	for carry != 0 {
		if tail.Next == nil {
			tail.Next = &listnode.ListNode{Val: 0}
		}
		carry += tail.Next.Val
		carry, tail.Next.Val, tail = carry/10, carry%10, tail.Next
	}

	return head
}
