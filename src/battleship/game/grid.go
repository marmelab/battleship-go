package game

// Grid is the board on which the battle ships are positionned
type Grid struct {
	Size  int
	Ships []Ship
	Cells []Cell
}

// NewGrid create a square grid of the given size
func NewGrid(gridSize int) Grid {
	cells := []Cell{}

	for row := 0; row < gridSize; row++ {
		for column := 0; column < gridSize; column++ {
			cells = append(cells, Cell{row, column})
		}
	}

	return Grid{gridSize, []Ship{}, cells}
}

// AddShip add a ship to the grid
func AddShip(grid Grid, ship Ship) Grid {
	grid.Ships = append(grid.Ships, ship)
	return grid
}
