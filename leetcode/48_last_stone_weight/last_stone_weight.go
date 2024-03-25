package last_stone_weight

import (
	"container/heap"
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n)
func LastStoneWeight(stones []int) int {
	stonesMaxHeap := MaxHeap(stones)
	heap.Init(&stonesMaxHeap)

	for stonesMaxHeap.Len() > 1 {
		y, x := heap.Pop(&stonesMaxHeap).(int), heap.Pop(&stonesMaxHeap).(int)

		if y != x {
			heap.Push(&stonesMaxHeap, y-x)
		}
	}

	if stonesMaxHeap.Len() == 1 {
		return stonesMaxHeap.Pop().(int)
	}

	return 0
}

type MaxHeap []int

func (maxHeap MaxHeap) Len() int {
	return len(maxHeap)
}

func (maxHeap MaxHeap) Less(i, j int) bool {
	return maxHeap[i] > maxHeap[j]
}

func (maxHeap MaxHeap) Swap(i, j int) {
	maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i]
}

func (maxHeap *MaxHeap) Push(x any) {
	*maxHeap = append(*maxHeap, x.(int))
}

func (maxHeap *MaxHeap) Pop() any {
	current := *maxHeap
	value := current[len(current)-1]
	*maxHeap = current[:len(current)-1]

	return value
}
