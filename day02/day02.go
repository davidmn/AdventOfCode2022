package day02

import (
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

type Result int64

const (
	Win  Result = 0
	Draw        = 1
	Lose        = 2
)

func CalculateScore(path string) int {
	scoreTable := map[Result]int{
		Win:  6,
		Draw: 3,
		Lose: 0,
	}

	myDecryptionTable := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	opponentDecryptionTable := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	choiceScores := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	score := 0

	lines := utils.ReadFile(path)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		splitLine := strings.Split(line, " ")

		opponentsChoice := opponentDecryptionTable[splitLine[0]]
		myChoice := myDecryptionTable[splitLine[1]]

		result := gameResult(myChoice, opponentsChoice)

		score += choiceScores[myChoice]
		score += scoreTable[result]
	}

	return score
}

func gameResult(myChoice string, opponentChoice string) Result {
	rules := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	if myChoice == opponentChoice {
		return Draw
	}

	if opponentChoice == rules[myChoice] {
		return Win
	} else {
		return Lose
	}
}
