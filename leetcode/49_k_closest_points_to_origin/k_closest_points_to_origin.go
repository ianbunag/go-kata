package k_closest_points_to_origin

import (
	"container/heap"
)

// Average time complexity: O(n log k)
// Worst time complexity:   O(n log k)
// Space complexity:        O(k)
func KClosest(points [][]int, k int) (result [][]int) {
	kClosestPoints := NewKClosestPoints(k)
	heap.Init(&kClosestPoints)

	for _, point := range points {
		kClosestPoints.Add(point[0], point[1])
	}
	for _, closestPoint := range kClosestPoints.points {
		result = append(result, []int{closestPoint.x, closestPoint.y})
	}

	return result
}

type ClosestPoint struct {
	x        int
	y        int
	distance int
}

type KClosestPoints struct {
	points []ClosestPoint
	k      int
}

func NewKClosestPoints(k int) KClosestPoints {
	return KClosestPoints{
		points: make([]ClosestPoint, 0),
		k:      k,
	}
}

// Average time complexity: O(log k)
// Worst time complexity:   O(log k)
// Space complexity:        O(1)
func (kClosestPoints *KClosestPoints) Add(x, y int) {
	heap.Push(kClosestPoints, ClosestPoint{
		x:        x,
		y:        y,
		distance: (x * x) + (y * y), // Formula derived from Euclidean distance, with square root and second point (0,0) omitted.
	})

	if kClosestPoints.Len() > kClosestPoints.k {
		heap.Pop(kClosestPoints)
	}
}

func (kClosestPoints KClosestPoints) Len() int {
	return len(kClosestPoints.points)
}

func (kClosestPoints KClosestPoints) Less(i, j int) bool {
	return kClosestPoints.points[i].distance > kClosestPoints.points[j].distance
}

func (kClosestPoints KClosestPoints) Swap(i, j int) {
	kClosestPoints.points[i], kClosestPoints.points[j] = kClosestPoints.points[j], kClosestPoints.points[i]
}

func (kClosestPoints *KClosestPoints) Push(x any) {
	kClosestPoints.points = append(kClosestPoints.points, x.(ClosestPoint))
}

func (kClosestPoints *KClosestPoints) Pop() any {
	current := kClosestPoints.points
	value := current[len(current)-1]
	kClosestPoints.points = current[:len(current)-1]

	return value
}
