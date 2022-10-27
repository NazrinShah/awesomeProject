package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Factorial_0(t *testing.T) {
	assert.Equal(t, 1, factorial(0))
}
