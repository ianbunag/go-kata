package funny_dots_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test25FunnyDots(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "25FunnyDots Suite")
}
