package main

import (
	"log"
	"strings"

	"github.com/salasberryfin/advent-of-code-2022/files"
)

/*
first column:
	A: rock
	B: paper
	C: scissors
second column:
	X: rock 		- 1
	Y: paper 		- 2
	Z: scissors 	- 3
score for round:
	0 lost
	3 draw
	6 win
*/

var (
	moveScore = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	winMove = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	drawMove = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	loseMove = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	expectedResult = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}
	roundScore = map[string]int{
		"lose": 0,
		"draw": 3,
		"win":  6,
	}
	//file = "test-input"
	file = "input"
)

func selectMove(opp, me string) string {
	if expectedResult[me] == "lose" {
		return loseMove[opp]
	} else if expectedResult[me] == "win" {
		return winMove[opp]
	}

	return drawMove[opp]
}

func main() {
	content := files.ReadFile(strings.Join([]string{"day2", file}, "-"))
	// use files.ReadFileByLine better: TBD
	splitContent := strings.Split(string(content), "\n")
	var score int
	for _, v := range splitContent {
		if len(v) == 0 {
			break
		}
		moves := strings.Split(v, " ")
		opponent, expected := moves[0], moves[1]
		me := selectMove(opponent, expected)
		score += moveScore[me]
		if winMove[opponent] == me {
			score += roundScore["win"]
		} else if drawMove[opponent] == me {
			score += roundScore["draw"]
		}
	}
	log.Println("Final score is: ", score)
}
