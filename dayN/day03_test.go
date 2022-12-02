package dayN

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayNPartN(t *testing.T) {
	score := DoTheThing("../inputs/day2.txt")
	assert.Equal(t, 12586, score)
}
