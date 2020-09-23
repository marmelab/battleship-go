package game

import (
	"battleship/game"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestGetNextBestShots(t *testing.T) {
	grid := game.NewGrid(3)
	grid = game.AddShoot(grid, game.Cell{1, 1})

	expected := []game.Cell{
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
	}

	actual := game.GetNextBestShots(game.Cell{1, 1}, grid)

	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Shoot at the center"))
}

func TestGetNextBestShotsOnBorders(t *testing.T) {
	grid := game.NewGrid(1)

	expected := []game.Cell{}

	actual := game.GetNextBestShots(game.Cell{0, 0}, grid)

	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Shoot at the center"))
}

func TestGetNextBestShotsWithObstacleAbove(t *testing.T) {
	grid := game.NewGrid(3)
	grid = game.AddShoot(grid, game.Cell{0, 1})

	expected := []game.Cell{
		{1, 0},
		{1, 2},
		{2, 1},
	}

	actual := game.GetNextBestShots(game.Cell{1, 1}, grid)

	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Shoot at the center"))
}

func TestGetNextBestShotsWithObstacleBelow(t *testing.T) {
	grid := game.NewGrid(3)
	grid = game.AddShoot(grid, game.Cell{2, 1})

	expected := []game.Cell{
		{0, 1},
		{1, 0},
		{1, 2},
	}

	actual := game.GetNextBestShots(game.Cell{1, 1}, grid)

	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Shoot at the center"))
}

func TestGetNextBestShotsWithMultipleObstacles(t *testing.T) {
	grid := game.NewGrid(4)
	grid = game.AddShoot(grid, game.Cell{1, 1})
	grid = game.AddShoot(grid, game.Cell{2, 1})

	expected := []game.Cell{
		{0, 0},
		{0, 2},
		{3, 1},
	}

	actual := game.GetNextBestShots(game.Cell{0, 1}, grid)

	then.AssertThat(t, actual, is.EqualTo(expected).Reason("Shoot at the center"))
}
