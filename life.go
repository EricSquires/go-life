package main

import (
	"time"
)

const (
	aliveChar string = "■"
	deadChar  string = "□"
)

func clear() {
	print("\033[H\033[2J")
}

func draw(board Board) {
	clear()

	for i := 0; i < board.Width; i++ {
		for j := 0; j < board.Height; j++ {
			if board.Cells[i][j].IsAlive {
				print(aliveChar)
			} else {
				print(deadChar)
			}
		}

		print("\n")
	}
}

func main() {
	pauseDuration := time.Millisecond * 500
	var board = InitBoard(50, 50)

	for i := 0; i < 60; i++ {
		draw(board)
		board = board.NextState()
		time.Sleep(pauseDuration)
	}
}
