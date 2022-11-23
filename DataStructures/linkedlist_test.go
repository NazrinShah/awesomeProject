package ds_test

import (
	"github.com/NazrinShah/awesomeProject/DataStructures/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_linked_list_insert_001(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)

	assert.Equal(t, 2, l.Count())
}

func Test_linked_list_remove_only_elem(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)

	v, err := l.Remove(0)

	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.Equal(t, 0, l.Count())
}

func Test_linked_list_remove_first_elem(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	v, err := l.Remove(0)

	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.Equal(t, 2, l.Count())
}

func Test_linked_list_remove_middle_elem(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	v, err := l.Remove(1)

	assert.Nil(t, err)
	assert.Equal(t, 2, v)
	assert.Equal(t, 2, l.Count())
}

func Test_linked_list_remove_last_elem(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	v, err := l.Remove(2)
	assert.Equal(t, 2, l.Count())

	assert.Nil(t, err)
	assert.Equal(t, 3, v)
}

func Test_linked_list_remove_invalid_001(t *testing.T) {
	l := pkg.LinkedList{}
	_ = l.Insert(1)
	_ = l.Insert(2)
	_ = l.Insert(3)

	_, err := l.Remove(10)
	assert.Equal(t, 3, l.Count())

	assert.NotNil(t, err)
}

func Test_linked_list_remove_invalid_002(t *testing.T) {
	l := pkg.LinkedList{}

	_, err := l.Remove(0)

	assert.NotNil(t, err)
	assert.Equal(t, 0, l.Count())
}
