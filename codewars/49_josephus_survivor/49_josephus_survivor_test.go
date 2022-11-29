package josephus_survivor_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "fmt"

	. "github.com/yvnbunag/go-kata/codewars/49_josephus_survivor"
)

var _ = Describe("49JosephusSurvivor", func() {
	It("should handle basic tests", func() {
		dotest(1, 1, 1)
		dotest(1, 2, 1)
		dotest(1, 3, 1)
		dotest(7, 3, 4)
		dotest(11, 19, 10)
		dotest(40, 3, 28)
		dotest(14, 2, 13)
		dotest(100, 1, 100)
	})
})

func dotest(n, k, e int) {
	Println(n, k)
	Expect(JosephusSurvivor(n, k)).To(Equal(e))
}
