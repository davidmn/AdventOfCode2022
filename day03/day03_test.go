package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	score := DoTheThing("../inputs/day2.txt")
	assert.Equal(t, 12586, score)
}
