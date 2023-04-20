package palindromic_substrings_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test30PalindromicSubstrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "30PalindromicSubstrings Suite")
}
