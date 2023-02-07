package lib_list_node_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestListNode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ListNode Suite")
}
