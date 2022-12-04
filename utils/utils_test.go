package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	lines := ReadFile("../inputs/day1.txt")
	assert.Equal(t, 2247, len(lines))
	assert.Equal(t, "", lines[8])
}

func TestIntersectionOneIntersection(t *testing.T) {
	a := []rune{'a', 'b'}
	b := []rune{'a'}
	intersection := Intersection(a, b)
	assert.Equal(t, 1, len(intersection))
	assert.Equal(t, 'a', intersection[0])
}

func TestIntersectionMany(t *testing.T) {
	a := []rune{'a', 'b', 'c'}
	b := []rune{'a', 'b', 'c'}
	intersection := Intersection(a, b)
	assert.Equal(t, 3, len(intersection))
}

func TestIntersectionNone(t *testing.T) {
	a := []rune{'a', 'b', 'c'}
	b := []rune{'d'}
	intersection := Intersection(a, b)
	assert.Equal(t, 0, len(intersection))
}

func TestIntersectionEmpty(t *testing.T) {
	a := []rune{'a', 'b', 'c'}
	b := []rune{}
	intersection := Intersection(a, b)
	assert.Equal(t, 0, len(intersection))
}

func TestIntersectionBothEmpty(t *testing.T) {
	a := []rune{}
	b := []rune{}
	intersection := Intersection(a, b)
	assert.Equal(t, 0, len(intersection))
}

func TestIntersectionWithInts(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2}
	intersection := Intersection(a, b)
	assert.Equal(t, 1, len(intersection))
	assert.Equal(t, 2, intersection[0])
}

func TestRemoveDuplicatesInts(t *testing.T) {
	a := []int{1, 2, 2, 3}

	b := RemoveDuplicates(a)

	assert.Equal(t, 3, len(b))
}

func TestRemoveDuplicatesRunes(t *testing.T) {
	a := []rune{'a', 'a'}

	b := RemoveDuplicates(a)

	assert.Equal(t, 1, len(b))
}
