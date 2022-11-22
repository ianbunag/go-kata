package tic_tac_toe_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test48TicTacToe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "48TicTacToe Suite")
}
