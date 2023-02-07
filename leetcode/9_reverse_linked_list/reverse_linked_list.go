package reverse_linked_list

import . "github.com/yvnbunag/go-kata/lib/list_node"

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	reversedHead := &ListNode{Val: head.Val, Next: nil}

	for head.Next != nil {
		reversedHead = &ListNode{Val: head.Next.Val, Next: reversedHead}
		head = head.Next
	}

	return reversedHead
}
