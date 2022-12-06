package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6Part1Example(t *testing.T) {
	marker := SolvePart1("../inputs/day6example.txt")
	assert.Equal(t, 11, marker)
}

func TestDay6Part1(t *testing.T) {
	marker := SolvePart1("../inputs/day6.txt")
	assert.Equal(t, 1658, marker)
}

func TestDay6Part2Example(t *testing.T) {
	marker := SolvePart2("../inputs/day6example.txt")
	assert.Equal(t, 26, marker)
}

func TestDay6Part2(t *testing.T) {
	marker := SolvePart2("../inputs/day6.txt")
	assert.Equal(t, 2260, marker)
}
