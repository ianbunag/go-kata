package snail_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/37_snail"
)

var _ = Describe("37Snail", func() {
	It("should resolve 0x0", func() {
		snailMap := [][]int{{}}
		Expect(Snail(snailMap)).To(Equal([]int{}))
	})

	It("should resolve 1x1", func() {
		snailMap := [][]int{{1}}
		Expect(Snail(snailMap)).To(Equal([]int{1}))
	})

	It("should resolve 2x2", func() {
		snailMap := [][]int{
			{1, 2},
			{4, 3},
		}
		Expect(Snail(snailMap)).To(Equal([]int{1, 2, 3, 4}))
	})

	It("should resolve 3x3", func() {
		snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		Expect(Snail(snailMap)).To(Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}))
	})

	It("should resolve 4x4", func() {
		snailMap := [][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		}
		Expect(Snail(snailMap)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}))
	})

	It("should resolve 5x5", func() {
		snailMap := [][]int{
			{1, 2, 3, 4, 5},
			{16, 17, 18, 19, 6},
			{15, 24, 25, 20, 7},
			{14, 23, 22, 21, 8},
			{13, 12, 11, 10, 9},
		}
		Expect(Snail(snailMap)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}))
	})
})
