package house_robber_2_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/29_house_robber_2"
)

var _ = Describe("29HouseRobber2", func() {
	It("should rob maximum amount", func() {
		Expect(Rob([]int{2, 3, 2})).To(Equal(3))
		Expect(Rob([]int{1, 2, 3, 1})).To(Equal(4))
		Expect(Rob([]int{1, 2, 3})).To(Equal(3))
		Expect(Rob([]int{1})).To(Equal(1))
	})
})
