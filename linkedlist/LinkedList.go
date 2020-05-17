package linkedlist

type LinkedList struct {
	head *node
	tail *node
	Size int
}

func (l *LinkedList) isEmpty() bool {
	return l.Size == 0
}

func (l LinkedList) String() string {
	if l.isEmpty() {
		return "List is empty"
	} else {
		return l.head.String()
	}
}

func BuildEmptyLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		Size: 0,
	}
}

func (l *LinkedList) Push(element int) *LinkedList {
	head := &node{
		value: element,
		next:  l.head,
	}
	l.head = head
	if l.tail == nil {
		l.tail = head
	}
	l.Size++
	return l
}

func (l *LinkedList) Append(element int) *LinkedList {
	if l.isEmpty() {
		return l.Push(element)
	} else {
		newTail := Node(element)
		l.tail.next = newTail
		l.tail = l.tail.next
		l.Size++
		return l
	}
}

func (l *LinkedList) NodeAt(index int) *node {
	currentIndex := 0
	currentHead := l.head

	for currentIndex < index && currentHead != nil {
		currentHead = currentHead.next
		currentIndex++
	}
	return currentHead
}

func (l *LinkedList) Insert(element int, afterNode *node) *node {
	if afterNode == l.tail {
		l.Append(element)
		return l.tail
	} else {
		newNode := &node{
			value: element,
			next: afterNode.next,
		}
		afterNode.next = newNode
		l.Size++
		return newNode
	}
}

func (l *LinkedList) Pop() int {
	if !l.isEmpty() {
		l.Size--
	}
	result := l.head.value
	l.head = l.head.next

	if l.isEmpty() {
		l.tail = nil
	}
	return result
}

func (l *LinkedList) RemoveLast() int {
	newHead := l.head

	if newHead.next == nil {
		return l.Pop()
	}
	l.Size--

	var prev = newHead
	var current = newHead
	var next = current.next

	for next != nil {
		prev = current
		current = next
		next = current.next
	}

	prev.next = nil
	l.tail = prev

	return current.value
}

func (l *LinkedList) RemoveAfter(n *node) int {
	result := n.next

	if n.next == l.tail {
		l.tail = n
	}
	if n.next != nil {
		l.Size--
	}
	n.next = n.next.next
	return result.value
}

func (l* LinkedList) Reverse() *LinkedList {
	current := l.head
	var previous *node = nil
	for current != nil {
		next := current.next
		current.next = previous
		previous = current
		if previous.next == nil {
			l.tail = previous
		}
		l.head = current
		current = next
	}
	return l
}
