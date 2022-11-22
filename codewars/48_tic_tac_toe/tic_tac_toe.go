package tic_tac_toe

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func IsSolved(rawBoard [3][3]int) int {
	board := NewBoard(rawBoard)
	unfinished := 0
	getLineStatuses := []GetLineStatus{
		func() LineStatus { return board.getRow(0).getStatus() },
		func() LineStatus { return board.getRow(1).getStatus() },
		func() LineStatus { return board.getRow(2).getStatus() },
		func() LineStatus { return board.getColumn(0).getStatus() },
		func() LineStatus { return board.getColumn(1).getStatus() },
		func() LineStatus { return board.getColumn(2).getStatus() },
		func() LineStatus { return board.getDiagonal(DirectionRightToLeft).getStatus() },
		func() LineStatus { return board.getDiagonal(DirectionLeftToRight).getStatus() },
	}

	for _, getLineStatus := range getLineStatuses {
		lineStatus := getLineStatus()

		if lineStatus > StatusDraw {
			return int(lineStatus)
		}

		if lineStatus == StatusUnfinished {
			unfinished += 1
		}
	}

	if unfinished > 0 {
		return int(StatusUnfinished)
	}

	return int(StatusDraw)
}

type LineStatus int

const (
	StatusUnfinished LineStatus = iota - 1
	StatusDraw
	StatusWinnerX
	StatusWinnerO
)

type GetLineStatus func() LineStatus

type Mark int

const (
	MarkNone Mark = iota
	MarkX
	MarkO
)

type Line [3]Mark

func NewLine(rawLine [3]int) Line {
	return Line{Mark(rawLine[0]), Mark(rawLine[1]), Mark(rawLine[2])}
}

func (line Line) isComplete() bool {
	for _, entry := range line {
		if entry == MarkNone {
			return false
		}
	}

	return true
}

func (line Line) getStatus() LineStatus {
	if !line.isComplete() {
		return StatusUnfinished
	}

	if line[0] != line[1] || line[1] != line[2] {
		return StatusDraw
	}

	if line[0] == MarkX {
		return StatusWinnerX
	}

	return StatusWinnerO
}

type Board [3]Line

type DiagonalDirection int

const (
	DirectionRightToLeft DiagonalDirection = iota
	DirectionLeftToRight
)

func NewBoard(rawBoard [3][3]int) Board {
	return Board{NewLine(rawBoard[0]), NewLine(rawBoard[1]), NewLine(rawBoard[2])}
}

func (board Board) getRow(row int) Line {
	return Line{board[row][0], board[row][1], board[row][2]}
}

func (board Board) getColumn(column int) Line {
	return Line{board[0][column], board[1][column], board[2][column]}
}

func (board Board) getDiagonal(direction DiagonalDirection) Line {
	if direction == DirectionRightToLeft {
		return Line{board[0][2], board[1][1], board[2][0]}
	}

	return Line{board[0][0], board[1][1], board[2][2]}
}
