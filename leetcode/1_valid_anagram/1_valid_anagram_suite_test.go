package valid_anagram_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1ValidAnagram(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1ValidAnagram Suite")
}
