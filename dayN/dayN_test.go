package dayN

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const day string = "N"

func TestDayNPart1Example(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%sexample.txt", day)
	result := SolvePart1(path)
	assert.Equal(t, 0, result)
}

func TestDayNPart1(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%s.txt", day)
	result := SolvePart1(path)
	assert.Equal(t, 0, result)
}

func TestDayNPart2Example(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%sexample.txt", day)
	result := SolvePart2(path)
	assert.Equal(t, 0, result)
}

func TestDayNPart2(t *testing.T) {
	path := fmt.Sprintf("../inputs/day%s.txt", day)
	result := SolvePart2(path)
	assert.Equal(t, 0, result)
}
