package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Factorial_Is_Zero(t *testing.T) {
	assert.Equal(t, 1, factorial(0))
}

func Test_Factorial_Is_Negative(t *testing.T) {
	assert.Equal(t, 1, factorial(-1))
}

func Test_Factorial_Is_Three(t *testing.T) {
	assert.Equal(t, 6, factorial(3))
}
