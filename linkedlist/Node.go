package linkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

func (n *node) hasNext() bool {
	return n.next != nil
}

func (n node) String() string {
	if n.hasNext() {
		return fmt.Sprintf("%v -> %v", n.value, n.next.String())
	} else {
		return fmt.Sprintf("%v", n.value)
	}
}

func Node(value int) *node {
	return &node{
		value: value,
		next:  nil,
	}
}
