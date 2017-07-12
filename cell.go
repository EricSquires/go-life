package main

// Cell represents a single cell in a game of life board
type Cell struct {
	X, Y    int
	IsAlive bool
}

// NextState returns the next state for the given cell
func (c Cell) NextState(cells [][]Cell) bool {
	var numNeighbors int

	for i := c.X - 1; i <= c.X+1; i++ {
		for j := c.Y - 1; j <= c.Y+1; j++ {
			var currentX = getIndex(i, len(cells))
			var currentY = getIndex(j, len(cells[currentX]))

			// Only increment if we're not looking at the current cell (that's not a neighbor no matter how lonely it is)
			if cells[currentX][currentY].IsAlive && !(currentX == c.X && currentY == c.Y) {
				numNeighbors++
			}
		}
	}

	// 3 neighbors always results in life, 2 only results in life if the cell was already alive to begin with
	return numNeighbors == 3 || (cells[c.X][c.Y].IsAlive && numNeighbors == 2)
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
