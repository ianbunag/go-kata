package plus_one_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/39_plus_one"
)

var _ = Describe("39PlusOne", func() {
	It("should add one", func() {
		Expect(PlusOne([]int{0})).To(Equal([]int{1}))
		Expect(PlusOne([]int{1})).To(Equal([]int{2}))
		Expect(PlusOne([]int{5})).To(Equal([]int{6}))
		Expect(PlusOne([]int{8})).To(Equal([]int{9}))
		Expect(PlusOne([]int{9})).To(Equal([]int{1, 0}))
		Expect(PlusOne([]int{1, 2, 3})).To(Equal([]int{1, 2, 4}))
		Expect(PlusOne([]int{4, 3, 2, 1})).To(Equal([]int{4, 3, 2, 2}))
		Expect(PlusOne([]int{9, 9, 9, 9})).To(Equal([]int{1, 0, 0, 0, 0}))
	})
})
