package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	board := [8]int{}
	solve(board[:], 0)
}

func IsSave(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]-i == col-row || board[i]+i == col+row {
			return false
		}
	}
	return true
}

func solve(board []int, row int) {
	if row == len(board) {
		for _, col := range board {
			z01.PrintRune(rune(col + '1'))
		}
		z01.PrintRune('\n')
		return
	}
	for col := 0; col < len(board); col++ {
		if IsSave(board, row, col) {
			board[row] = col
			solve(board, row+1)
			board[row] = -1
		}
	}
}
