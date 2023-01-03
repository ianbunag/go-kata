package group_anagrams_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2GroupAnagrams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2GroupAnagrams Suite")
}
