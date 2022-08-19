package gold_mine_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/geek-for-geeks/4_gold_mine"
)

var _ = Describe("4GoldMine", func() {
	It("should find maximum amount", func() {
		Expect(FindMaximumAmount([][]int{})).To(Equal(0))
		Expect(FindMaximumAmount([][]int{
			{0, 0},
		})).To(Equal(0))
		Expect(FindMaximumAmount([][]int{
			{1, 0},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{0, 1},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{1, 1},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{0, 0},
			{0, 0},
		})).To(Equal(0))
		Expect(FindMaximumAmount([][]int{
			{0, 1},
			{0, 0},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{0, 0},
			{0, 1},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{1, 0},
			{0, 0},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{0, 0},
			{1, 0},
		})).To(Equal(1))
		Expect(FindMaximumAmount([][]int{
			{1, 0},
			{0, 1},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{1, 1},
			{0, 0},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{0, 0},
			{1, 1},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{0, 1},
			{1, 0},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{2},
			{1},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{1},
			{2},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{2},
			{2},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{1, 1},
			{1, 1},
		})).To(Equal(2))
		Expect(FindMaximumAmount([][]int{
			{1, 3, 3},
			{2, 1, 4},
			{0, 6, 4},
		})).To(Equal(12))
		Expect(FindMaximumAmount([][]int{
			{1, 3, 1, 5},
			{2, 2, 4, 1},
			{5, 0, 2, 3},
			{0, 6, 1, 2},
		})).To(Equal(16))
		Expect(FindMaximumAmount([][]int{
			{10, 33, 13, 15},
			{22, 21, 04, 1},
			{5, 0, 2, 3},
			{0, 6, 14, 12},
		})).To(Equal(83))
		Expect(FindMaximumAmount([][]int{
			{7, 0, 0, 0, 7},
			{0, 0, 0, 0, 0},
			{0, 0, 21, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{6, 6, 6, 6, 6},
		})).To(Equal(35))
		Expect(FindMaximumAmount([][]int{
			{5, 0, 0, 0, 100},
			{0, 6, 0, 0, 0},
			{0, 0, 7, 0, 0},
			{0, 0, 0, 8, 0},
			{0, 0, 0, 0, 9},
		})).To(Equal(118))
	})
})
