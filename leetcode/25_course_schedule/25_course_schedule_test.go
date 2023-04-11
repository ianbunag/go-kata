package course_schedule

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("25CourseSchedule", func() {
	It("should return true on feasible schedule", func() {
		Expect(CanFinish(2, [][]int{{1, 0}})).To(Equal(true))
		Expect(CanFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})).To(Equal(true))
	})

	It("should return false on simple circular schedule", func() {
		Expect(CanFinish(1, [][]int{{0, 0}})).To(Equal(false))
		Expect(CanFinish(2, [][]int{{1, 0}, {0, 1}})).To(Equal(false))
	})

	It("should return false on complex circular schedule", func() {
		Expect(CanFinish(3, [][]int{{0, 1}, {1, 2}, {2, 0}})).To(Equal(false))
		Expect(CanFinish(4, [][]int{{0, 1}, {2, 3}, {1, 2}, {3, 0}})).To(Equal(false))
		Expect(CanFinish(3, [][]int{{1, 0}, {2, 0}, {0, 2}})).To(Equal(false))
	})
})
