package kth_largest_element_in_an_array

import "container/heap"

// Average time complexity: O(n log k)
// Worst time complexity:   O(n log k)
// Space complexity:        O(n)
func FindKthLargest(nums []int, k int) int {
	maxHeap := MaxHeap(nums)
	heap.Init(&maxHeap)

	for ; k > 1; k -= 1 {
		heap.Pop(&maxHeap)
	}

	return heap.Pop(&maxHeap).(int)
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
