package main

import (
	"math/rand"
	"time"
)

// InitBoard sets up the initial state for a blank life board
func InitBoard(width, height int) [][]bool {
	rand.Seed(time.Now().Unix())

	var cells = make([][]bool, width)

	for i := 0; i < width; i++ {
		cells[i] = make([]bool, height)

		for j := 0; j < height; j++ {
			if rand.Float32() <= 0.3 {
				cells[i][j] = true
			}
		}
	}

	return cells
}

// BoardState returns the next state of all cells in the board
func BoardState(board [][]bool) [][]bool {
	nextBoard := make([][]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			nextBoard[i][j] = CellState(board, i, j)
		}
	}

	return nextBoard
}
