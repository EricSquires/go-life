package main

// CellState returns the next state for the given cell
func CellState(board [][]bool, x, y int) bool {
	var numNeighbors int

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			var currentX = getIndex(i, len(board))
			var currentY = getIndex(j, len(board[i]))

			if board[currentX][currentY] {
				numNeighbors++
			}
		}
	}

	// 3 neighbors always results in life, 2 only results in life if the cell was already alive to begin with
	return numNeighbors == 3 || (board[x][y] && numNeighbors == 2)
}

func getIndex(n, limit int) int {
	if n < 0 {
		return n + limit
	} else if n >= limit {
		return n - limit
	} else {
		return n
	}
}
