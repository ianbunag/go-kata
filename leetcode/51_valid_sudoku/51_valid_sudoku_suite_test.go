package valid_sudoku_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test51ValidSudoku(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "51ValidSudoku Suite")
}
