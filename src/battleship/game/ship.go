package game

// Ship represents a battle ship on a grid
type Ship struct {
	Length int
	Cells  []Cell
}

// ShipsOverlapHorizontally tells if a ship overlaps another one on a particular cell in a row
func ShipsOverlapHorizontally(ship Ship, gridShip Ship, cell Cell) bool {
	for _, gridShipCell := range gridShip.Cells {
		for i := 0; i < ship.Length; i++ {
			cellToCheck := Cell{cell.Row, cell.Column + i}
			if cellToCheck == gridShipCell {
				return true
			}
		}
	}

	return false
}

// ShipsOverlapVertically tells if a ship overlaps another one on a particular cell in a column
func ShipsOverlapVertically(ship Ship, gridShip Ship, cell Cell) bool {
	for _, gridShipCell := range gridShip.Cells {
		for i := 0; i < ship.Length; i++ {
			cellToCheck := Cell{cell.Row + i, cell.Column}
			if cellToCheck == gridShipCell {
				return true
			}
		}

	}

	return false
}
