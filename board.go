package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
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
			cells[i][j] = Cell{i, j, false}
		}
	}

	return board
}

// InitBoardRandom sets up a Board with each cell having the given chance to have its starting state set to alive
func InitBoardRandom(width, height int, chanceAlive float32) Board {
	rand.Seed(time.Now().Unix())

	var cells = make([][]Cell, width)
	board := Board{width, height, cells}

	for i := 0; i < width; i++ {
		cells[i] = make([]Cell, height)

		for j := 0; j < height; j++ {
			cells[i][j] = Cell{i, j, rand.Float32() <= chanceAlive}
		}
	}

	return board
}

// LoadFromFile parses a file and creates an identical Board based on the file
func LoadFromFile(path string) Board {
	boardData, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	boardString := strings.Split(strings.TrimSpace(string(boardData)), "\n")

	width := strings.TrimSpace(boardString[0])
	board := Board{len(width), len(boardString), make([][]Cell, len(width))}

	for i := 0; i < board.Width; i++ {
		board.Cells[i] = make([]Cell, board.Height)

		for j := 0; j < board.Height; j++ {
			var isAlive bool
			currentChar := string(boardString[i][j])

			if strings.TrimSpace(currentChar) == "" {
				continue
			}

			if currentChar == AliveChar || currentChar == AliveAltChar {
				isAlive = true
			} else if currentChar != DeadChar && currentChar != DeadAltChar {
				panic("Unrecognized character found in board definition: '" + currentChar + "'")
			}

			board.Cells[i][j] = Cell{i, j, isAlive}
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
