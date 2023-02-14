package linked_list_cycle

import (
	. "github.com/yvnbunag/go-kata/lib/list_node"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func HasCycle(head *ListNode) bool {
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer, fastPointer = slowPointer.Next, fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
