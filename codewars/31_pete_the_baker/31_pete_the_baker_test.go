package pete_the_baker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/31_pete_the_baker"
)

var _ = Describe("31PeteTheBaker", func() {
	It("Sample tests", func() {
		Expect(Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})).To(Equal(2))
		Expect(Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000})).To(Equal(0))
		Expect(Cakes(map[string]int{"cocoa": 20, "oil": 52}, map[string]int{"cocoa": 2017, "oil": 47})).To(Equal(0))
	})
})
