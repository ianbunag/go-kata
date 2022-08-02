package beeramid_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/16_beeramid"
)

var _ = Describe("16Beeramid", func() {
	It("Sample tests", func() {
		Expect(Beeramid(9, 2.0)).To(Equal(1))
		Expect(Beeramid(21, 1.5)).To(Equal(3))
		Expect(Beeramid(-1, 4.0)).To(Equal(0))
	})
})
