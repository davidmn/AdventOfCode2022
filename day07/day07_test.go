package day07

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day string = "7"

func TestDay7Part1Example(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%sexample.txt", day)
	result := SolvePart1(path)
	assert.Equal(t, 95437, result)
}

func TestDay7Part1(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%s.txt", day)
	result := SolvePart1(path)
	assert.Equal(t, 1886043, result)
}

func TestDay7Part2Example(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%sexample.txt", day)
	result := SolvePart2(path)
	assert.Equal(t, 24933642, result)
}

func TestDay7Part2(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%s.txt", day)
	result := SolvePart2(path)
	assert.Equal(t, 3842121, result)
}
