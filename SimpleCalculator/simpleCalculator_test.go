package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Division_By_Zero(t *testing.T) {
	_, err := divide(0, 0)
	assert.NotNil(t, err)

	assert.Equal(t, "denominator cannot be 0", err.Error())
}
