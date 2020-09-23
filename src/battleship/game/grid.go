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
