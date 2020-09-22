package game

import (
	"battleship/scoreboard"
)

// Grid is the board on which the battle ships are positionned
type Grid struct {
	Size  int
	Ships []Ship
	Cells []Cell
}

// NewGrid create a square grid of the given size
func NewGrid(gridSize int) (Grid, error) {
	cells := []Cell{}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {
			cells = append(cells, Cell{row, column})
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

			shipCanBePlacedHorizontally := shipCanBePlacedHorizontally(Cell{row, column}, ship, grid)
			if shipCanBePlacedHorizontally {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row][column+l]++
				}
			}

			shipCanBePlacedVertically := shipCanBePlacedVertically(Cell{row, column}, ship, grid)
			if shipCanBePlacedVertically {
				for l := 0; l < ship.Length; l++ {
					scoreBoard.Cells[row+l][column]++
				}
			}
		}
	}

	return &scoreBoard, nil
}

func shipCanBePlacedHorizontally(cell Cell, ship Ship, grid Grid) bool {
	if cell.Column+ship.Length > grid.Size {
		return false
	}

	for _, gridShip := range grid.Ships {
		for _, gridShipCell := range gridShip.Cells {
			for i := 0; i < ship.Length; i++ {
				cellToCheck := Cell{cell.Row, cell.Column + i}
				if cellToCheck == gridShipCell {
					return false
				}
			}

		}
	}

	return true
}

func shipCanBePlacedVertically(cell Cell, ship Ship, grid Grid) bool {
	if cell.Row+ship.Length > grid.Size {
		return false
	}

	for _, gridShip := range grid.Ships {
		for _, gridShipCell := range gridShip.Cells {
			for i := 0; i < ship.Length; i++ {
				cellToCheck := Cell{cell.Row + i, cell.Column}
				if cellToCheck == gridShipCell {
					return false
				}
			}

		}
	}

	return true
}
