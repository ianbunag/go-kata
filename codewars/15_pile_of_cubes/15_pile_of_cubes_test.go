package pile_of_cubes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/15_pile_of_cubes"
)

func dotest(n int, exp int) {
	var ans = FindNb(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("15PileOfCubes", func() {
	It("should handle basic cases", func() {
		dotest(4183059834009, 2022)
		dotest(24723578342962, -1)
	})
})
