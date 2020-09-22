package game

import (
	"battleship/game"
	"battleship/grid"
	"battleship/scoreboard"
	"fmt"
	"strconv"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestGetScoreBoardWithOneCellLongShip(t *testing.T) {
	// Given one 1 cell long ship
	ship := game.Ship{1}

	cells := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	// on a 3x3 grid
	actual, _ := grid.GetScoreBoard(3, ship) // adresse récupérée

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("1 cell long ship on 3x3 grid"))
	DisplayScoreBoard(actual, ship)
}

func TestGetScoreBoardWithTwoCellsLongShip(t *testing.T) {
	// Given one 2 cells long ship
	ship := game.Ship{2}

	cells := [][]int{
		{2, 3, 2},
		{3, 4, 3},
		{2, 3, 2},
	}

	expected := &scoreboard.ScoreBoard{3, cells}

	// When computing its possible positions
	// on a 3x3 grid
	actual, _ := grid.GetScoreBoard(3, ship)

	// Then it should equals this score board
	then.AssertThat(t, actual, is.EqualTo(expected).Reason("2 cells long ship on 3x3 grid"))
	DisplayScoreBoard(actual, ship)
}

func TestGetScoreBoardShouldNotBeComputableWithTooLongShip(t *testing.T) {
	// Given one 4 cells long ship
	ship := game.Ship{4}

	// When computing its possible positions
	// on a 3x3 grid
	_, error := grid.GetScoreBoard(3, ship)

	// Then it should not be computed and return an error
	if error == nil {
		t.Error("The ship is too long. Its length must be inferior to the grid size.")
	}
}

func DisplayScoreBoard(scoreBoard *scoreboard.ScoreBoard, ship game.Ship) {
	fmt.Println(scoreboard.GetStringFromScoreBoard(scoreBoard), "  "+strconv.Itoa(ship.Length)+" long ship on 3x3 grid")
}
