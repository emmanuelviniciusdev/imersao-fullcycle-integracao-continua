package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(1, 1.5)
	assert.Equal(t, 9999999, result)
}
