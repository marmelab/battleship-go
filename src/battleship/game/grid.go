package game

// Grid is the board on which the battle ships are positionned
type Grid struct {
	Size   int
	Ships  []Ship
	Shoots []Cell
}

// NewGrid create a square grid of the given size
func NewGrid(gridSize int) Grid {
	return Grid{gridSize, []Ship{}, []Cell{}}
}

// AddShip add a ship to the grid
func AddShip(grid Grid, ship Ship) Grid {
	grid.Ships = append(grid.Ships, ship)
	return grid
}

// AddShoot add a shoot to the grid
func AddShoot(grid Grid, shoot Cell) Grid {
	grid.Shoots = append(grid.Shoots, shoot)
	return grid
}

// GetNextBestShots returns the list of possible next shots given the previous one
func GetNextBestShots(lastShot Cell, grid Grid) []Cell {
	nextShots := []Cell{}

	nextShots = huntTopCell(lastShot, grid, nextShots)
	nextShots = huntLeftCell(lastShot, grid, nextShots)
	nextShots = huntRightCell(lastShot, grid, nextShots)
	nextShots = huntBottomCell(lastShot, grid, nextShots)

	return nextShots
}

func huntTopCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Row <= 0 {
		return nextShots
	}

	topCell := Cell{cell.Row - 1, cell.Column}

	if hasShoot(topCell, grid) {
		return huntTopCell(topCell, grid, nextShots)
	}

	return append(nextShots, topCell)
}

func huntLeftCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Column <= 0 {
		return nextShots
	}

	leftCell := Cell{cell.Row, cell.Column - 1}

	if hasShoot(leftCell, grid) {
		return huntLeftCell(leftCell, grid, nextShots)
	}

	return append(nextShots, leftCell)
}

func huntRightCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Column >= grid.Size-1 {
		return nextShots
	}

	rightCell := Cell{cell.Row, cell.Column + 1}

	if hasShoot(rightCell, grid) {
		return huntRightCell(rightCell, grid, nextShots)
	}

	return append(nextShots, rightCell)
}

func huntBottomCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Row >= grid.Size-1 {
		return nextShots
	}

	bottomCell := Cell{cell.Row + 1, cell.Column}

	if hasShoot(bottomCell, grid) {
		return huntBottomCell(bottomCell, grid, nextShots)
	}

	return append(nextShots, bottomCell)
}

func hasShoot(cell Cell, grid Grid) bool {
	for _, shoot := range grid.Shoots {
		if shoot == cell {
			return true
		}
	}

	return false
}
