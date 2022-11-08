package make_deadfish_swim_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/13_make_deadfish_swim"
)

var _ = Describe("13MakeDeadfishSwim", func() {
	It("just o", func() {
		Expect(Parse("ooo")).To(Equal([]int{0, 0, 0}))
	})
	It("o&i", func() {
		Expect(Parse("ioioio")).To(Equal([]int{1, 2, 3}))
	})
	It("o&i&d", func() {
		Expect(Parse("idoiido")).To(Equal([]int{0, 1}))
	})
	It("o&i&d&s", func() {
		Expect(Parse("isoisoiso")).To(Equal([]int{1, 4, 25}))
	})
	It("codewars", func() {
		Expect(Parse("codewars")).To(Equal([]int{0}))
	})
})
