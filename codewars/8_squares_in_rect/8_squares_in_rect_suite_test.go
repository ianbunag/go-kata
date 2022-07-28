package squares_in_rect_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test8SquaresInRect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "8SquaresInRect Suite")
}
