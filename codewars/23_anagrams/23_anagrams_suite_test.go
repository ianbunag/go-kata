package anagrams_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test23Anagrams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "23Anagrams Suite")
}
