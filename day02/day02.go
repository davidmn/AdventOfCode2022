package day02

import (
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

type Result int64

const (
	Win  Result = 0
	Draw Result = 1
	Lose Result = 2
)

func CalculateScore(path string) int {
	score := 0

	lines := utils.ReadFile(path)

	for i := 0; i < len(lines); i++ {
		score += calculateRoundScore(lines[i])
	}

	return score
}

func CalculateScorePart2(path string) int {
	score := 0

	lines := utils.ReadFile(path)

	for i := 0; i < len(lines); i++ {
		score += calculateRoundScorePart2(lines[i])
	}

	return score
}

func calculateRoundScore(line string) int {
	scoreTable := getScoreTable()
	myDecryptionTable := getMyDecryptionTable()
	opponentDecryptionTable := getOpponentDecryptionTable()
	choiceScores := getChoiceScores()

	splitLine := strings.Split(line, " ")

	opponentsChoice := opponentDecryptionTable[splitLine[0]]
	myChoice := myDecryptionTable[splitLine[1]]

	result := gameResult(myChoice, opponentsChoice)

	return choiceScores[myChoice] + scoreTable[result]
}

func calculateRoundScorePart2(line string) int {
	scoreTable := getScoreTable()
	opponentDecryptionTable := getOpponentDecryptionTable()
	choiceScores := getChoiceScores()
	roundTargetResultTable := getRoundTargetResult()

	splitLine := strings.Split(line, " ")

	opponentsChoice := opponentDecryptionTable[splitLine[0]]
	roundTarget := roundTargetResultTable[splitLine[1]]

	myChoice := calculateTargetChoice(roundTarget, opponentsChoice)

	result := gameResult(myChoice, opponentsChoice)

	return choiceScores[myChoice] + scoreTable[result]
}

func calculateTargetChoice(target Result, opponentsChoice string) string {
	rules := getRules()

	switch target {
	case Draw:
		return opponentsChoice
	case Lose:
		return rules[opponentsChoice]
	default:
		return rules[rules[opponentsChoice]]
	}
}

func gameResult(myChoice string, opponentChoice string) Result {
	rules := getRules()

	if myChoice == opponentChoice {
		return Draw
	}

	if opponentChoice == rules[myChoice] {
		return Win
	} else {
		return Lose
	}
}

func getScoreTable() map[Result]int {
	return map[Result]int{
		Win:  6,
		Draw: 3,
		Lose: 0,
	}
}

func getMyDecryptionTable() map[string]string {
	return map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
}

func getOpponentDecryptionTable() map[string]string {
	return map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}
}

func getChoiceScores() map[string]int {
	return map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
}

func getRoundTargetResult() map[string]Result {
	return map[string]Result{
		"Z": Win,
		"Y": Draw,
		"X": Lose,
	}
}

func getRules() map[string]string {
	return map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}
}
