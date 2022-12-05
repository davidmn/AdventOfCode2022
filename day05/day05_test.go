package day05

import (
	"testing"

	"github.com/davidmn/AdventOfCode2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay5Part1Example(t *testing.T) {
	output := SolvePart1("../inputs/day5example.txt")
	assert.Equal(t, "CMZ", output)
}

func TestDay5Part1(t *testing.T) {
	output := SolvePart1("../inputs/day5.txt")
	assert.Equal(t, "GRTSWNJHH", output)
}

func TestDay5Part2Example(t *testing.T) {
	output := SolvePart2("../inputs/day5example.txt")
	assert.Equal(t, "MCD", output)
}

func TestDay5Part2(t *testing.T) {
	output := SolvePart2("../inputs/day5.txt")
	assert.Equal(t, "QLFQDBBHM", output)
}

func TestFindHowManyStacks(t *testing.T) {
	exampleLines := utils.ReadFile("../inputs/day5example.txt")
	numberOfStacks := findHowManyStacks(exampleLines)
	assert.Equal(t, 3, numberOfStacks)

	lines := utils.ReadFile("../inputs/day5.txt")
	numberOfStacks = findHowManyStacks(lines)
	assert.Equal(t, 9, numberOfStacks)
}
