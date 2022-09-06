package directions_reduction_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/codewars/24_directions_reduction"
)

func dotest(arr []string, exp []string) {
	var ans = DirReduc(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("24DirectionsReduction", func() {
	It("should handle basic cases", func() {
		var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
		dotest(a, []string{})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
		dotest(a, []string{"WEST"})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
		dotest(a, []string{"NORTH"})
		a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
		dotest(a, []string{"NORTH", "WEST", "SOUTH", "EAST"})
	})
})
