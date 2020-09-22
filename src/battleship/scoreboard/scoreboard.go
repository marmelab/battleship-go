package scoreboard

import "strconv"

type ScoreBoard struct {
	Length  int
	Squares [][]int
}

func MakeScoreBoard(size int) ScoreBoard {

	squares := make([][]int, size)

	for i := 0; i < size; i++ {
		squares[i] = make([]int, size)
	}

	return ScoreBoard{size, squares}
}

func GetStringFromScoreBoard(scoreBoard *ScoreBoard) string {
	res := ""
	for row := 0; row < len(scoreBoard.Squares); row++ {
		for column := 0; column < len(scoreBoard.Squares); column++ {
			res += strconv.Itoa(scoreBoard.Squares[row][column])

			if column < len(scoreBoard.Squares)-1 {
				res += " "
			} else if row < len(scoreBoard.Squares)-1 {
				res += "\n"
			}
		}
	}

	return res
}
