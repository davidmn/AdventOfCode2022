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
	intersection1 := intersection([]rune(line1), []rune(line2))
	intersection2 := intersection([]rune(line2), []rune(line3))
	intersectionOfIntersections := intersection(intersection1, intersection2)

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

	intersection := intersection(firstPocket, secondPocket)

	if len(intersection) == 1 {
		return intersection[0], nil
	}

	return 0, errors.New("oh no")
}

func intersection(slice1, slice2 []rune) (inter []rune) {
	hash := make(map[rune]bool)
	for _, e := range slice1 {
		hash[e] = true
	}
	for _, e := range slice2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDuplicates(inter)
	return
}

func removeDuplicates(elements []rune) (deduplicated []rune) {
	encountered := make(map[rune]bool)
	for _, element := range elements {
		if !encountered[element] {
			deduplicated = append(deduplicated, element)
			encountered[element] = true
		}
	}
	return
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
