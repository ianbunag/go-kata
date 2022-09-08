package spin_words_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test26SpinWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "26SpinWords Suite")
}
