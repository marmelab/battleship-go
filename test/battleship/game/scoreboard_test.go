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
	ship := game.Ship{1}
	grid, _ := game.NewGrid(3)

	cells := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	actual, _ := game.GetScoreBoard(grid, ship) // adresse récupérée

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("1 cell long ship on 3x3 grid"))
	displayScoreBoard(actual, ship)
}

func TestGetScoreBoardWithTwoCellsLongShip(t *testing.T) {
	// Given one 2 cells long ship
	ship := game.Ship{2}
	grid, _ := game.NewGrid(3)

	cells := [][]int{
		{2, 3, 2},
		{3, 4, 3},
		{2, 3, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	// on a 3x3 grid
	actual, _ := game.GetScoreBoard(grid, ship)

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("2 cells long ship on 3x3 grid"))
	displayScoreBoard(actual, ship)
}

func TestGetScoreBoardShouldNotBeComputableWithTooLongShip(t *testing.T) {
	// Given one 4 cells long ship
	// on a 3x3 grid
	ship := game.Ship{4}
	grid, _ := game.NewGrid(3)

	cells := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	actual, _ := game.GetScoreBoard(grid, ship)

	// Then it should not be computed
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Too long ship on 3x3 grid"))
	displayScoreBoard(actual, ship)
}

func displayScoreBoard(scoreBoard *scoreboard.ScoreBoard, ship game.Ship) {
	fmt.Println(scoreboard.ToString(scoreBoard), "  "+strconv.Itoa(ship.Length)+" long ship on 3x3 grid")
}
