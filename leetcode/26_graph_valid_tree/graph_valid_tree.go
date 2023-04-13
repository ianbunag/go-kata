package graph_valid_tree

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func ValidTree(count int, edges [][]int) bool {
	adjacencyList := make([][]int, count)
	visited := make([]bool, count)
	disconnected := count

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		adjacencyList[node1] = append(adjacencyList[node1], node2)
	}

	for node := 0; node < count; node += 1 {
		if len(adjacencyList[node]) == 0 {
			continue
		}

		nodes := []int{node}

		for len(nodes) > 0 {
			currentNode := nodes[0]

			if visited[currentNode] {
				return false
			}

			nodes = append(nodes[1:], adjacencyList[currentNode]...)
			// Remove the current node from the graph by clearing its adjacency list,
			// so that it is not revisited during subsequent iterations of the BFS
			// traversal.
			adjacencyList[currentNode] = []int{}
			visited[currentNode] = true
			disconnected -= 1
		}
	}

	return disconnected == 0
}
