package surrounded_regions_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/yvnbunag/go-kata/leetcode/62_surrounded_regions"
)

var _ = Describe("62SurroundedRegions", func() {
	It("should capture regions", func() {
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		expectedBoard := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		Solve(board)
		Expect(board).To(Equal(expectedBoard))

		board = [][]byte{
			{'X'},
		}
		expectedBoard = [][]byte{
			{'X'},
		}
		Solve(board)
		Expect(board).To(Equal(expectedBoard))

		board = [][]byte{
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
		}
		expectedBoard = [][]byte{
			{'X', 'O', 'X', 'O', 'X', 'O'},
			{'O', 'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'X', 'O', 'X', 'O', 'X'},
		}
		Solve(board)
		Expect(board).To(Equal(expectedBoard))

		board = [][]byte{
			{'O', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'X', 'X', 'X'},
		}
		expectedBoard = [][]byte{
			{'O', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'X', 'X', 'X'},
		}
		Solve(board)
		Expect(board).To(Equal(expectedBoard))
	})
})
