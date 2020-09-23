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

func GetNextBestShots(lastShot Cell, grid Grid) []Cell {
	nextShots := []Cell{}

	nextShots = HuntTopCell(lastShot, grid, nextShots)
	nextShots = HuntLeftCell(lastShot, grid, nextShots)
	nextShots = HuntRightCell(lastShot, grid, nextShots)
	nextShots = HuntBottomCell(lastShot, grid, nextShots)

	return nextShots
}

func HuntTopCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Row <= 0 {
		return nextShots
	}

	topCell := Cell{cell.Row - 1, cell.Column}

	if HasShoot(topCell, grid) {
		return HuntTopCell(topCell, grid, nextShots)
	} else {
		return append(nextShots, topCell)
	}
}

func HuntLeftCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Column <= 0 {
		return nextShots
	}

	leftCell := Cell{cell.Row, cell.Column - 1}

	if HasShoot(leftCell, grid) {
		return HuntLeftCell(leftCell, grid, nextShots)
	} else {
		return append(nextShots, leftCell)
	}
}

func HuntRightCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Column >= grid.Size-1 {
		return nextShots
	}

	rightCell := Cell{cell.Row, cell.Column + 1}

	if HasShoot(rightCell, grid) {
		return HuntRightCell(rightCell, grid, nextShots)
	} else {
		return append(nextShots, rightCell)
	}
}

func HuntBottomCell(cell Cell, grid Grid, nextShots []Cell) []Cell {
	if cell.Row >= grid.Size-1 {
		return nextShots
	}

	bottomCell := Cell{cell.Row + 1, cell.Column}

	if HasShoot(bottomCell, grid) {
		return HuntBottomCell(bottomCell, grid, nextShots)
	} else {
		return append(nextShots, bottomCell)
	}
}

func HasShoot(cell Cell, grid Grid) bool {
	for _, shoot := range grid.Shoots {
		if shoot == cell {
			return true
		}
	}

	return false
}
