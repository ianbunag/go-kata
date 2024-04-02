package redundant_connection_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test55RedundantConnection(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "55RedundantConnection Suite")
}
