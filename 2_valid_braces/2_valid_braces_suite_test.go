package valid_braces_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2ValidBraces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2ValidBraces Suite")
}
