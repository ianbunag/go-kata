package reverse_words_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test1ReverseWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1ReverseWords Suite")
}
