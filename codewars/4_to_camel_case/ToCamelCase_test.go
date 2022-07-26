package to_camel_case_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/codewars/4_to_camel_case"
)

func dotest(str, exp string) {
	fmt.Println("input:", str)
	var ans = ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("ToCamelCase", func() {

	It("should handle basic cases", func() {
		dotest("", "")
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
		dotest("the-Stealth-Warrior", "theStealthWarrior")
		dotest("the-Stealth-warrior", "theStealthWarrior")
	})
})
