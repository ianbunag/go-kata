package redundant_connection

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
// Guide: 									https://youtu.be/FXWRE67PLL0?feature=shared
func FindRedundantConnection(edges [][]int) []int {
	parents, ranks := make([]int, len(edges)+1), make([]int, len(edges)+1)

	for index := range parents {
		parents[index], ranks[index] = index, 1
	}

	find := func(node int) int {
		parent := parents[node]

		for parent != parents[parent] {
			parents[parent] = parents[parents[parent]]
			parent = parents[parent]
		}

		return parent
	}
	union := func(a, b int) bool {
		aParent, bParent := find(a), find(b)

		if aParent == bParent {
			return false
		}

		if ranks[aParent] >= ranks[bParent] {
			ranks[aParent] += ranks[bParent]
			parents[bParent] = aParent
		} else {
			ranks[bParent] += ranks[aParent]
			parents[aParent] = bParent
		}

		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return edges[len(edges)-1]
}

// Worst time complexity:   O(n^2)
// Average time complexity: O(n^2)
// Space complexity:        O(n)
func FindRedundantConnectionV1(edges [][]int) (redundant []int) {
	nodes := make([]Node, len(edges))

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		nodes[a].value, nodes[b].value = a, b

		if nodes[a].Connected(nodes[b]) {
			redundant = edge
		}

		nodes[a].Connect(&nodes[b])
	}

	return redundant
}

type Node struct {
	value       int
	connections []*Node
}

func (node Node) Connected(connection Node) bool {
	visited := map[int]bool{}
	queue := append([]*Node{}, node.connections...)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.value] {
			continue
		}

		if current.value == connection.value {
			return true
		}

		queue = append(queue, current.connections...)
		visited[current.value] = true
	}

	return false
}

func (node *Node) Connect(connection *Node) {
	node.connections = append(node.connections, connection)
	connection.connections = append(connection.connections, node)
}
