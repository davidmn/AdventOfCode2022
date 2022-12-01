package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCaloriesDay1Part1(t *testing.T) {
	var result = CountCaloriesDay1Part1()
	assert.Equal(t, 69177, result)
}

func TestCountCaloriesDay1Part2(t *testing.T) {
	var result = CountCaloriesDay1Part2()
	assert.Equal(t, 207456, result)
}
