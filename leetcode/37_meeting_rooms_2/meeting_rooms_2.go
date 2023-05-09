package meeting_rooms_2

import (
	"container/heap"
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n)
func MinMeetingRooms(intervals []*Interval) (minRooms int) {
	rooms := 0
	startIntervals := make(IntHeap, len(intervals))
	endIntervals := make(IntHeap, len(intervals))

	for index, interval := range intervals {
		startIntervals[index], endIntervals[index] = interval.Start, interval.End
	}

	heap.Init(&startIntervals)
	heap.Init(&endIntervals)

	for startIntervals.Len() > 0 {
		if startIntervals.Peek() < endIntervals.Peek() {
			heap.Pop(&startIntervals)
			rooms += 1
			minRooms = max(minRooms, rooms)
			continue
		}

		heap.Pop(&endIntervals)
		rooms -= 1
	}

	return
}

type Interval struct {
	Start, End int
}

type IntHeap []int

func (intHeap IntHeap) Len() int {
	return len(intHeap)
}

func (intHeap IntHeap) Less(i, j int) bool {
	return intHeap[i] < intHeap[j]
}

func (intHeap IntHeap) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}

func (intHeap IntHeap) Peek() int {
	return (intHeap)[0]
}

func (intHeap *IntHeap) Push(x any) {
	*intHeap = append(*intHeap, x.(int))
}

func (intHeap *IntHeap) Pop() any {
	current := *intHeap
	currentLength := len(current)
	*intHeap = current[0 : currentLength-1]

	return current[currentLength-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
