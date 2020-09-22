package scoreboard

import "strconv"

// ScoreBoard holds the probabilities of finding ships on the cells
type ScoreBoard struct {
	Length int
	Cells  [][]int
}

// NewScoreBoard creates a square board of the given size
func NewScoreBoard(size int) ScoreBoard {

	cells := make([][]int, size)

	for i := 0; i < size; i++ {
		cells[i] = make([]int, size)
	}

	return ScoreBoard{size, cells}
}

// ToString returns returns the stringify score board
func ToString(scoreBoard *ScoreBoard) string {
	res := ""
	for row := 0; row < len(scoreBoard.Cells); row++ {
		for column := 0; column < len(scoreBoard.Cells); column++ {
			res += strconv.Itoa(scoreBoard.Cells[row][column])

			if column < len(scoreBoard.Cells)-1 {
				res += " "
			} else if row < len(scoreBoard.Cells)-1 {
				res += "\n"
			}
		}
	}

	return res
}
