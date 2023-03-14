package word_search_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/19_word_search"
)

var _ = Describe("19WordSearch", func() {
	It("should search if simple word exists in board", func() {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		Expect(Exist(board, "ABCCED")).To(Equal(true))
		Expect(Exist(board, "SEE")).To(Equal(true))
		Expect(Exist(board, "ABCB")).To(Equal(false))
	})

	It("should search if complex word exists in board", func() {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'E', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		Expect(Exist(board, "ABCESEEEFS")).To(Equal(true))
		Expect(Exist(board, "ABCEFSADEESE")).To(Equal(true))
	})
})
