package main

import (
	"bufio"
	"os"
)

func Main1D() int {
	var board [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 8 && scanner.Scan(); i++ {
		line := scanner.Text()
		board = append(board, []rune(line))
	}
	return Solve1D(board)
}

func Solve1D(board [][]rune) int {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				takeStep(i, j, 0, -1, board)
				takeStep(i, j, 0, 1, board)
				takeStep(i, j, -1, 0, board)
				takeStep(i, j, 1, 0, board)
			} else if board[i][j] == 'B' {
				takeStep(i, j, -1, -1, board)
				takeStep(i, j, -1, 1, board)
				takeStep(i, j, 1, -1, board)
				takeStep(i, j, 1, 1, board)
			}
		}
	}

	cnt := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' || board[i][j] == 'B' || board[i][j] == '+' {
				cnt++
			}
		}
	}
	return 64 - cnt
}

func takeStep(row, col, stepRow, stepCol int, board [][]rune) {
	for row, col = row+stepRow, col+stepCol; 0 <= row && 0 <= col && 7 >= row && 7 >= col && board[row][col] != 'R' && board[row][col] != 'B'; row, col = row+stepRow, col+stepCol {
		board[row][col] = '+'
	}
}
