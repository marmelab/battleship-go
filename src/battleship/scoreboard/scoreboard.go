package scoreboard

import "strconv"

func GetStringFromScoreBoard(scoreBoard [][]int) string {
	res := ""
	for row := 0; row < len(scoreBoard); row++ {
		for column := 0; column < len(scoreBoard); column++ {
			res += strconv.Itoa(scoreBoard[row][column])

			if column < len(scoreBoard)-1 {
				res += " "
			} else if row < len(scoreBoard)-1 {
				res += "\n"
			}
		}
	}

	return res
}
