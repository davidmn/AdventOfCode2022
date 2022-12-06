package day06

import (
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

func SolvePart1(path string) int {
	file := utils.ReadFile(path)

	return findStartOfPacket(file[0], 4)
}

func SolvePart2(path string) int {
	file := utils.ReadFile(path)

	return findStartOfPacket(file[0], 14)
}

func findStartOfPacket(stream string, packetSize int) int {
	frontIndex := packetSize
	backIndex := 0

	for i := 0; i < len(stream); i++ {
		window := stream[backIndex:frontIndex]

		unique := true

		for i := 0; i < len(window); i++ {
			isCharacterUnique := strings.Count(window, string(window[i])) == 1
			unique = unique && isCharacterUnique
		}

		if unique {
			return frontIndex
		}

		frontIndex++
		backIndex++
	}

	return -1
}
