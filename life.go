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

func draw(board [][]bool) {
	clear()

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] {
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
		time.Sleep(pauseDuration)
	}
}
