package valid_braces_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/yvnbunag/go-kata/2_valid_braces"
)

func singleTest(str string, res bool) {
	It(fmt.Sprintf("should return %v for \"%v\"", res, str), func() {
		Expect(ValidBraces(str)).To(Equal(res))
	})
}

var _ = Describe("Valid Braces", func() {
	singleTest("(){}[]", true)
	singleTest("([{}])", true)
	singleTest("(}", false)
	singleTest("[(])", false)
	singleTest("[({)](]", false)
	singleTest("((", false)
	singleTest("(())]", false)
})
