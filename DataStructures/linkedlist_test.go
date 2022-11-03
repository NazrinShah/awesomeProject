package ds_test

import (
	"awesomeProject/DataStructures/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_linked_list_insert_001(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)

	assert.Equal(t, 2, l.Count())
}

func Test_linked_list_remove_001(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)

	v, err := l.Remove(0)

	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}

func Test_linked_list_remove_002(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	v, err := l.Remove(1)

	assert.Nil(t, err)
	assert.Equal(t, 2, v)
}

func Test_linked_list_remove_003(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	v, err := l.Remove(2)

	assert.Nil(t, err)
	assert.Equal(t, 3, v)
}

func Test_linked_list_remove_004(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	_, err := l.Remove(10)

	assert.NotNil(t, err)
}
