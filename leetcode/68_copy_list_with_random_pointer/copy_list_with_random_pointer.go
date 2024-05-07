package copy_list_with_random_pointer

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
func CopyRandomList(head *Node) *Node {
	copies := map[*Node]*Node{}

	for node := head; node != nil; node = node.Next {
		copies[node] = &Node{Val: node.Val}
	}

	for node := head; node != nil; node = node.Next {
		copies[node].Next, copies[node].Random = copies[node.Next], copies[node.Random]
	}

	return copies[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
