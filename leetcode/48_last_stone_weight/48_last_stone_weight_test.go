package last_stone_weight_test

import (
	"container/heap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/48_last_stone_weight"
)

var _ = Describe("48LastStoneWeight", func() {
	It("should return last stone weight", func() {
		Expect(LastStoneWeight([]int{2, 7, 4, 1, 8, 1})).To(Equal(1))
		Expect(LastStoneWeight([]int{2, 2})).To(Equal(0))
		Expect(LastStoneWeight([]int{2, 2, 2})).To(Equal(2))
		Expect(LastStoneWeight([]int{1, 3})).To(Equal(2))
		Expect(LastStoneWeight([]int{3, 7, 2})).To(Equal(2))
		Expect(LastStoneWeight([]int{10, 4, 2, 10})).To(Equal(2))
	})
})

var _ = Describe("MaxHeap", func() {
	It("should Pop the maximum value", func() {
		maxHeap := MaxHeap([]int{1, 10, 5})
		heap.Init(&maxHeap)

		Expect(heap.Pop(&maxHeap)).To(Equal(10))
		Expect(heap.Pop(&maxHeap)).To(Equal(5))
		heap.Push(&maxHeap, 4)
		heap.Push(&maxHeap, 9)
		Expect(heap.Pop(&maxHeap)).To(Equal(9))
		Expect(heap.Pop(&maxHeap)).To(Equal(4))
		Expect(heap.Pop(&maxHeap)).To(Equal(1))
	})
})
