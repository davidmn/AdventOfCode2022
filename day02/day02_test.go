package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	score := CalculateScore("../inputs/day2.txt")
	assert.Equal(t, 12586, score)
}

func TestDay2Part1Example(t *testing.T) {
	score := CalculateScore("../inputs/day2example.txt")
	assert.Equal(t, 15, score)
}

func TestDay2Part2(t *testing.T) {
	score := CalculateScorePart2("../inputs/day2.txt")
	assert.Equal(t, 13193, score)
}

func TestDay2Part2Example(t *testing.T) {
	score := CalculateScorePart2("../inputs/day2example.txt")
	assert.Equal(t, 12, score)
}
