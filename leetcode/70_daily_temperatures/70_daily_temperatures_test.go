package daily_temperatures_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/70_daily_temperatures"
)

var _ = Describe("70DailyTemperatures", func() {
	It("should return days to wait until warmer temperatures", func() {
		Expect(DailyTemperatures([]int{30})).To(Equal([]int{0}))
		Expect(DailyTemperatures([]int{30, 60, 90})).To(Equal([]int{1, 1, 0}))
		Expect(DailyTemperatures([]int{30, 40, 50, 60})).To(Equal([]int{1, 1, 1, 0}))
		Expect(DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})).To(Equal([]int{1, 1, 4, 2, 1, 1, 0, 0}))
	})
})
