package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1Example(t *testing.T) {
	score := Solve("../inputs/day3example.txt")
	assert.Equal(t, 157, score)
}

func TestDay2Part1(t *testing.T) {
	score := Solve("../inputs/day3.txt")
	assert.Equal(t, 8243, score)
}
