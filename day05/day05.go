package day05

import (
	//"fmt"
	"strconv"
	"strings"

	"github.com/davidmn/AdventOfCode2022/utils"
)

func SolvePart1(path string) string {
	lines := utils.ReadFile(path)

	stack := parseStacks(lines)
	output := runCrane(stack, lines, true)

	return output
}

func SolvePart2(path string) string {
	lines := utils.ReadFile(path)

	stack := parseStacks(lines)
	output := runCrane(stack, lines, false)

	return output
}

func runCrane(stacks [][]string, lines []string, moveIndividually bool) string {
	for i := 0; i < len(lines); i++ {
		isCommandLine := strings.HasPrefix(lines[i], "move")
		if !isCommandLine {
			continue
		}

		split := strings.Split(lines[i], " ")
		n, _ := strconv.Atoi(split[1])
		source, _ := strconv.Atoi(split[3])
		destination, _ := strconv.Atoi(split[5])
		source--
		destination--

		if moveIndividually {
			for i := 0; i < n; i++ {
				topCrateIndex := len(stacks[source]) - 1
				crate := stacks[source][topCrateIndex]
				stacks[source] = utils.RemoveIndex(stacks[source], topCrateIndex)
				stacks[destination] = append(stacks[destination], crate)
			}
		} else {
			topCrateIndex := len(stacks[source])
			crates := stacks[source][topCrateIndex-n:]
			for i := 0; i < n; i++ {
				topCrateIndex = len(stacks[source]) - 1
				stacks[source] = utils.RemoveIndex(stacks[source], topCrateIndex)
			}
			stacks[destination] = append(stacks[destination], crates...)
		}
	}

	topCrates := make([]string, 0)
	for i := 0; i < len(stacks); i++ {
		topCrates = append(topCrates, stacks[i][len(stacks[i])-1])
	}

	return strings.Join(topCrates, "")
}

func parseStacks(lines []string) [][]string {
	numberOfStacks := findHowManyStacks(lines)
	stacks := make([][]string, numberOfStacks)

	for i := 0; i < numberOfStacks; i++ {
		stacks[i] = make([]string, 0)
	}

	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], " 1") {
			break
		}
		splitLine := strings.Split(lines[i], "")

		stackCounter := 0
		for i := 1; i < len(splitLine); i += 4 {
			currentCrate := splitLine[i]
			if currentCrate != " " {
				stacks[stackCounter] = append(stacks[stackCounter], currentCrate)
			}
			stackCounter++
		}

	}

	for i := 0; i < numberOfStacks; i++ {
		utils.Reverse(stacks[i])
	}

	return stacks
}

func findHowManyStacks(lines []string) int {
	if len(lines) == 2 {
		return 0
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			stackDefinitions := strings.Split(lines[i-1], " ")

			maxStackNumber := 0
			for i := 0; i < len(stackDefinitions); i++ {
				if stackDefinitions[i] == "" {
					continue
				}
				maxStackNumber, _ = strconv.Atoi(stackDefinitions[i])
			}

			return maxStackNumber
		}
	}

	return 0
}
