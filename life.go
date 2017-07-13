package main

import (
	"os"
	"strconv"
	"time"
)

const (
	// AliveChar is the character used to denote an alive cell
	AliveChar string = "■"
	// DeadChar is the character used to denote a dead cell
	DeadChar string = "□"
	// AliveAltChar is an alternative alive character to use to make creating board files easier
	AliveAltChar string = "A"
	// DeadAltChar is an alternative dead character to use to make creating board files easier
	DeadAltChar string = "-"

	defaultSize = 50
)

func clear() {
	print("\033[H\033[2J")
}

func draw(board Board) {
	clear()

	for i := 0; i < board.Width; i++ {
		for j := 0; j < board.Height; j++ {
			if board.Cells[i][j].IsAlive {
				print(AliveChar)
			} else {
				print(DeadChar)
			}
		}

		print("\n")
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func initBoardFromArgs(args []string) Board {
	if len(args) <= 1 {
		return InitBoard(defaultSize, defaultSize)
	}

	var board Board
	size, err := strconv.Atoi(args[1])

	if err != nil {
		if fileExists(args[1]) {
			board = LoadFromFile(args[1])
		}
	} else if len(args) >= 3 {
		chanceAlive, err := strconv.ParseFloat(args[2], 32)

		if err == nil {
			board = InitBoardRandom(size, size, float32(chanceAlive))
		}
	}

	if board.Width == 0 && board.Height == 0 {
		if size == 0 {
			board = InitBoard(defaultSize, defaultSize)
		} else {
			board = InitBoard(size, size)
		}
	}

	return board
}

func main() {

	board := initBoardFromArgs(os.Args)
	pauseDuration := time.Millisecond * 500

	for i := 0; i < 60; i++ {
		draw(board)
		board = board.NextState()
		time.Sleep(pauseDuration)
	}
}
