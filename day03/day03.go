package day03

import (
	"errors"

	"github.com/davidmn/AdventOfCode2022/utils"
)

func Solve(path string) int {
	lines := utils.ReadFile(path)
	total := 0

	for i := 0; i < len(lines); i++ {
		letter, _ := processLine(lines[i])
		total += determinePriority(letter)
	}

	return total
}

func SolvePart2(path string) int {
	lines := utils.ReadFile(path)
	total := 0

	if len(lines)%3 != 0 {
		panic("wow now")
	}

	for i := 0; i < len(lines); i += 3 {
		letter, _ := processTrio(lines[i], lines[i+1], lines[i+2])
		total += determinePriority(letter)
	}

	return total
}

func processTrio(line1 string, line2 string, line3 string) (rune, error) {
	intersection1 := utils.Intersection([]rune(line1), []rune(line2))
	intersection2 := utils.Intersection([]rune(line2), []rune(line3))
	intersectionOfIntersections := utils.Intersection(intersection1, intersection2)

	if len(intersectionOfIntersections) == 1 {
		return intersectionOfIntersections[0], nil
	}
	return 0, errors.New("oh no")
}

func processLine(line string) (rune, error) {
	runeArray := []rune(line)
	halfLength := len(runeArray) / 2

	firstPocket := runeArray[0:halfLength]
	secondPocket := runeArray[halfLength:]

	intersection := utils.Intersection(firstPocket, secondPocket)

	if len(intersection) == 1 {
		return intersection[0], nil
	}

	return 0, errors.New("oh no")
}

func determinePriority(letter rune) int {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, item := range characters {
		if item == letter {
			return i + 1
		}
	}

	return 0
}
