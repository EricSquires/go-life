package main

import (
	"math/rand"
	"time"
)

// Board represents a game of life board
type Board struct {
	Width, Height int
	Cells         [][]Cell
}

// InitBoard sets up the initial state for a blank life board
func InitBoard(width, height int) Board {
	rand.Seed(time.Now().Unix())

	var cells = make([][]Cell, width)
	board := Board{width, height, cells}

	for i := 0; i < width; i++ {
		cells[i] = make([]Cell, height)

		for j := 0; j < height; j++ {
			cells[i][j] = Cell{i, j, rand.Float32() <= 0.3}
		}
	}

	return board
}

// NextState returns the next state of all cells in the board
func (b Board) NextState() Board {
	nextBoard := Board{b.Width, b.Height, make([][]Cell, b.Width)}

	for i := 0; i < nextBoard.Width; i++ {
		nextBoard.Cells[i] = make([]Cell, b.Height)

		for j := 0; j < nextBoard.Height; j++ {
			nextBoard.Cells[i][j] = Cell{i, j, b.Cells[i][j].NextState(b.Cells)}
		}
	}

	return nextBoard
}
