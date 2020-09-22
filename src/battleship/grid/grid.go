package grid

import (
	"battleship/game"
	"battleship/scoreboard"
	"errors"
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

func GetScoreBoard(gridSize int, ship game.Ship) (*scoreboard.ScoreBoard, error) {
	if ship.Length > gridSize {
		return nil, errors.New("The ship is too long. Its length must be inferior to the grid size.")
	}

	scoreBoard := scoreboard.MakeScoreBoard(gridSize)

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {

			if column+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Squares[row][column+l]++
				}
			}

			if row+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Squares[row+l][column]++
				}
			}
		}
	}

	return &scoreBoard, nil
}
