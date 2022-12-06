package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4Part1Example(t *testing.T) {
	score := SolvePart1("../inputs/day4example.txt")
	assert.Equal(t, 2, score)
}

func TestDay4Part1(t *testing.T) {
	score := SolvePart1("../inputs/day4.txt")
	assert.Equal(t, 657, score)
}

func TestDay4Part2Example(t *testing.T) {
	score := SolvePart2("../inputs/day4example.txt")
	assert.Equal(t, 4, score)
}

func TestDay4Part2(t *testing.T) {
	score := SolvePart2("../inputs/day4.txt")
	assert.Equal(t, 938, score)
}
