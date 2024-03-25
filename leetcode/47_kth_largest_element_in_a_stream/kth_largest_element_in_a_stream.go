package kth_largest_element_in_a_stream

import (
	"container/heap"
)

type KthLargest struct {
	minHeap []int
	k       int
}

// Average time complexity: O(n log k)
// Worst time complexity:   O(n log k)
// Space complexity:        O(n)
func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{k: k}
	heap.Init(&kthLargest)

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(n)
func (kthLargest *KthLargest) Add(val int) int {
	heap.Push(kthLargest, val)

	if kthLargest.Len() > kthLargest.k {
		heap.Pop(kthLargest)
	}

	return kthLargest.minHeap[0]
}

func (kthLargest KthLargest) Len() int {
	return len(kthLargest.minHeap)
}

func (kthLargest KthLargest) Less(i, j int) bool {
	return kthLargest.minHeap[i] < kthLargest.minHeap[j]
}

func (kthLargest KthLargest) Swap(i, j int) {
	kthLargest.minHeap[i], kthLargest.minHeap[j] = kthLargest.minHeap[j], kthLargest.minHeap[i]
}

func (kthLargest *KthLargest) Push(x interface{}) {
	kthLargest.minHeap = append(kthLargest.minHeap, x.(int))
}

func (kthLargest *KthLargest) Pop() interface{} {
	current := kthLargest.minHeap
	value := current[len(current)-1]
	kthLargest.minHeap = current[:len(current)-1]

	return value
}
