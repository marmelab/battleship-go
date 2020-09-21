package grid

import (
	"battleship/game"
)

type Grid struct {
	Size    int
	Squares []game.Coordinates
}

func NewGrid(gridSize int) (Grid, error) {
	squares := []game.Coordinates{}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {
			squares = append(squares, game.Coordinates{row, column})
		}
	}

	return Grid{gridSize, squares}, nil
}

func GetScoreBoard(gridSize int, ship game.Ship) [][]int {
	scoreBoard := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		scoreBoard[i] = make([]int, gridSize)
	}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {

			if column+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard[row][column+l]++
				}
			}

			if row+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard[row+l][column]++
				}
			}
		}
	}

	return scoreBoard
}
