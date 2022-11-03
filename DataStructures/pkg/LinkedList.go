package pkg

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *node
	size int
}

type node struct {
	value int
	next  *node
}

func (l *LinkedList) Count() int {
	return l.size
}

func (l *LinkedList) Insert(val int) error {

	newNode := &node{
		value: val,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
	} else {
		t := l.head

		for t.next != nil {
			t = t.next
		}

		t.next = newNode
	}

	l.size++
	return nil
}

func (l *LinkedList) Remove(i int) (int, error) {
	if i >= l.size || i < 0 || l.head == nil {
		return 0, errors.New(fmt.Sprintf("invalid index: %d provided but linkedlist has size %d", i, l.size))
	}

	t := l.head
	j := 0
	tVal := 0
	var err error
	err = nil

	// remove first node
	if i == 0 {
		tVal = l.head.value
		l.head = l.head.next
		l.size--
	} else {
		found := false

		// remove middle node
		for t.next.next != nil {
			if i == j+1 {
				tVal = t.next.value
				t.next = t.next.next
				l.size--
				found = true
				break
			}

			j++
			t = t.next
		}

		if !found {
			// remove last node
			tVal = t.next.value
			t.next = nil
			l.size--
		}
	}

	return tVal, err
}

func (l *LinkedList) Print() {
	t := l.head

	for t != nil {
		fmt.Println(t.value)
		t = t.next
	}
}

func (l *LinkedList) Retrieve(i int) (int, error) {

	if i < 0 || i >= l.size || l.head == nil {
		return 0, errors.New(fmt.Sprintf("invalid index: %d provided but linkedlist has size %d", i, l.size))
	}

	t := l.head

	for i > 0 {
		t = t.next
	}

	return t.value, nil
}
