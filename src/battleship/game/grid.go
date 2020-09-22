package game

import (
	"battleship/scoreboard"
)

// Grid is the board on which the battle ships are positionned
type Grid struct {
	Size  int
	Ships []Ship
	Cells []Coordinates
}

// NewGrid create a square grid of the given size
func NewGrid(gridSize int) (Grid, error) {
	cells := []Coordinates{}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {
			cells = append(cells, Coordinates{row, column})
		}
	}

	return Grid{gridSize, []Ship{}, cells}, nil
}

// AddShip add a ship to the grid
func AddShip(grid Grid, ship Ship) Grid {
	grid.Ships = append(grid.Ships, ship)
	return grid
}

// GetScoreBoard returns the score board of a ship positioned on a grid of the given size
func GetScoreBoard(grid Grid, ship Ship) (*scoreboard.ScoreBoard, error) {
	scoreBoard := scoreboard.NewScoreBoard(grid.Size)

	if ship.Length > grid.Size {
		return &scoreBoard, nil
	}

	for row := 0; row < grid.Size; row++ {
		for column := 0; column < grid.Size; column++ {

			if column+ship.Length <= grid.Size {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row][column+l]++
				}
			}

			if row+ship.Length <= grid.Size {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row+l][column]++
				}
			}
		}
	}

	return &scoreBoard, nil
}
