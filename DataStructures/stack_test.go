package ds_test

import (
	"awesomeProject/DataStructures/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stack_insert_001(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	v, _ := s.Peek()
	assert.Equal(t, 1, v)
}

func Test_stack_pop_001(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	v, _ := s.Pop()
	assert.Equal(t, 1, v)
}
