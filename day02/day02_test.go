package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	score := CalculateScore("../inputs/day2.txt")
	assert.Equal(t, 12586, score)
}

func TestDay2Example(t *testing.T) {
	score := CalculateScore("../inputs/day2example.txt")
	assert.Equal(t, 15, score)
}
