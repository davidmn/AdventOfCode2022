package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	var result = doThing(1)
	assert.Equal(t, 1, result, "doThing should return the input")
}
