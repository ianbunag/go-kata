package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/lib"
)

var _ = Describe("GenerateString", func() {
	It("should generate string with seed length less than length", func() {
		Expect(GenerateString("az", 1)).To(Equal("a"))
	})

	It("should generate string with seed length equal to length", func() {
		Expect(GenerateString("az", 2)).To(Equal("az"))
	})

	It("should generate string with seed length greater than length", func() {
		Expect(GenerateString("az", 3)).To(Equal("aza"))
	})
})
