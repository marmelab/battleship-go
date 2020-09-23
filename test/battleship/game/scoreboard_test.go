package game

import (
	"battleship/game"
	"battleship/scoreboard"
	"fmt"
	"strconv"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestGetScoreBoardWithOneCellLongShip(t *testing.T) {
	// Given one 1 cell long ship
	// on a 3x3 grid
	ship := game.Ship{1, []game.Cell{}}
	grid := game.NewGrid(3)

	cells := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	actual := scoreboard.GetScoreBoard(grid, ship) // adresse récupérée

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("1 cell long ship on 3x3 grid"))
	displayScoreBoard(actual, ship, grid)
}

func TestGetScoreBoardWithTwoCellsLongShip(t *testing.T) {
	// Given one 2 cells long ship
	ship := game.Ship{2, []game.Cell{}}
	grid := game.NewGrid(3)

	cells := [][]int{
		{2, 3, 2},
		{3, 4, 3},
		{2, 3, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	// on a 3x3 grid
	actual := scoreboard.GetScoreBoard(grid, ship)

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("2 cells long ship on 3x3 grid"))
	displayScoreBoard(actual, ship, grid)
}

func TestGetScoreBoardWithTooLongShip(t *testing.T) {
	// Given one 4 cells long ship
	// on a 3x3 grid
	ship := game.Ship{4, []game.Cell{}}
	grid := game.NewGrid(3)

	cells := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	actual := scoreboard.GetScoreBoard(grid, ship)

	// Then it should not be computed
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Too long ship on 3x3 grid"))
	displayScoreBoard(actual, ship, grid)
}

func TestGetScoreBoardWithObstacle(t *testing.T) {
	// Given a grid with a 1 cell long ship on it (the obstacle)
	// and considering a 2 cells long ship
	grid := game.NewGrid(3)
	ship := game.Ship{1, []game.Cell{{1, 2}}}
	grid = game.AddShip(grid, ship)

	computedShip := game.Ship{2, []game.Cell{}}

	cells := [][]int{
		{2, 3, 1},
		{3, 3, 0},
		{2, 3, 1},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing possible positions of the second ship
	actual := scoreboard.GetScoreBoard(grid, computedShip)

	// Then the resulting score board should equals the expected one
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("There is an obstacle on cell 1:2"))
	displayScoreBoard(actual, computedShip, grid)
}

func displayScoreBoard(scoreBoard *scoreboard.ScoreBoard, ship game.Ship, grid game.Grid) {
	message := scoreboard.ToString(scoreBoard)
	message += "  "
	message += strconv.Itoa(ship.Length) + " long ship on 3x3 grid"
	if len(grid.Ships) > 0 {
		message += " with obstacle"
	}
	fmt.Println(message)
}
