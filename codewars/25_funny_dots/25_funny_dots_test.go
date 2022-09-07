package funny_dots_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/25_funny_dots"
)

func dotest(n, m int, expected string) {
	Expect(Dot(n, m)).To(Equal(expected), "With n = %d, m = %d", n, m)
}

var _ = Describe("25FunnyDots", func() {
	It("Sample tests", func() {
		dotest(1, 1, "+---+\n| o |\n+---+")
		dotest(3, 2, "+---+---+---+\n| o | o | o |\n+---+---+---+\n| o | o | o |\n+---+---+---+")
	})
})
