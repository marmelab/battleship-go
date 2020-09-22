package grid

import (
	"battleship/game"
	"battleship/scoreboard"
)

// Grid is the board on which the battle ships are positionned
type Grid struct {
	Size  int
	Cells []game.Coordinates
}

// NewGrid create a square grid of the given size
func NewGrid(gridSize int) (Grid, error) {
	cells := []game.Coordinates{}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {
			cells = append(cells, game.Coordinates{row, column})
		}
	}

	return Grid{gridSize, cells}, nil
}

// GetScoreBoard returns the score board of a ship positioned on a grid of the given size
func GetScoreBoard(gridSize int, ship game.Ship) (*scoreboard.ScoreBoard, error) {
	scoreBoard := scoreboard.NewScoreBoard(gridSize)

	if ship.Length > gridSize {
		return &scoreBoard, nil
	}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {

			if column+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row][column+l]++
				}
			}

			if row+ship.Length <= gridSize {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row+l][column]++
				}
			}
		}
	}

	return &scoreBoard, nil
}
