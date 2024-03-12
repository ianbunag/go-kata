package binary_search_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test42BinarySearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "42BinarySearch Suite")
}
