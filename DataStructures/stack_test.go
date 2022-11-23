package ds_test

import (
	"github.com/NazrinShah/awesomeProject/DataStructures/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stack_insert_001(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	v, _ := s.Peek()
	assert.Equal(t, 1, v)
	assert.Equal(t, 1, s.Count())
}

func Test_stack_insert_002(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, _ := s.Peek()
	assert.Equal(t, 3, v)
	assert.Equal(t, 3, s.Count())
}

func Test_stack_pop_001(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	v, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, s.Count())
}

func Test_stack_pop_002(t *testing.T) {
	s := pkg.Stack{}
	s.Push(1)
	s.Push(2)
	v, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
	assert.Equal(t, 1, s.Count())
}

func Test_stack_pop_empty(t *testing.T) {
	s := pkg.Stack{}
	_, err := s.Pop()
	assert.NotNil(t, err)
	assert.Equal(t, 0, s.Count())
}
