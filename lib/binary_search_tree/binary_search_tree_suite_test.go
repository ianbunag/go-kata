package lib_binary_search_tree_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBinarySearchTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinarySearchTree Suite")
}
