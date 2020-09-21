package game

import (
	"battleship/game"
	"battleship/grid"
	"battleship/scoreboard"
	"fmt"
	"testing"
)

func TestScoreBoardWithOneShip(t *testing.T) {
	// Given one 2 long ship
	ship := game.Ship{2}

	// When computing its possible positions
	// on a 3x3 grid
	scoreBoard := grid.GetScoreBoard(3, ship)

	// Then it should equals this score board
	expected := "2 3 2\n3 4 3\n2 3 2"

	actual := scoreboard.GetStringFromScoreBoard(scoreBoard)

	if actual != expected {
		t.Errorf("Score board incorrect, got: %s, expected: %s.", actual, expected)
	} else {
		DisplayExpectedAndActualScoreBoards(expected, actual)
	}
}

func DisplayExpectedAndActualScoreBoards(expected string, actual string) {
	fmt.Println(expected, "  Expected")
	fmt.Println()
	fmt.Println(actual, "  Actual")
}
