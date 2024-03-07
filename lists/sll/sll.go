package sll

type List[T comparable] struct {
	head *node[T]
	tail *node[T]
	len  int
}

func New[T comparable]() *List[T] {
	return &List[T]{}
}

func (list *List[T]) Len() int {
	return list.len
}

func (list *List[T]) Push(val T) *List[T] {
	newNode := &node[T]{val: val}

	list.len++

	if list.head == nil {
		list.head = newNode
		list.tail = list.head

		return list
	}

	list.tail.next = newNode
	list.tail = newNode

	return list
}

func (list *List[T]) PopFront() (v T, ok bool) {
	if list.head == nil {
		return zeroValue[T](), false
	}

	oldHead := list.head

	list.head = oldHead.next

	list.len--

	if list.len == 0 {
		list.tail = nil
	}

	return oldHead.val, true
}

func (list *List[T]) Pop() (v T, ok bool) {
	if list.tail == nil {
		return zeroValue[T](), false
	}

	current := list.head
	newTail := current

	for current.next != nil {
		newTail = current
		current = current.next
	}

	list.tail = newTail
	list.tail.next = nil

	list.len--

	if list.len == 0 {
		list.head = nil
		list.tail = nil
	}

	return current.val, true
}

func (list *List[T]) PushFront(val T) *List[T] {
	newNode := &node[T]{val: val}

	list.len++

	if list.len == 0 {
		list.head = newNode
		list.tail = list.head

		return list
	}

	newNode.next = list.head
	list.head = newNode

	return list
}

func (list *List[T]) findByIndex(idx int) (n *node[T], ok bool) {
	if idx < 0 || idx >= list.len {
		return &node[T]{}, false
	}

	current := list.head

	for i := 0; i < idx; i++ {
		current = current.next
	}

	return current, true
}

func (list *List[T]) FindByIndex(idx int) (v T, ok bool) {
	n, found := list.findByIndex(idx)

	if !found {
		return zeroValue[T](), false
	}

	return n.val, true
}

func (list *List[T]) Insert(val T, idx int) bool {
	if idx < 0 || idx > list.len {
		return false
	}

	if idx == 0 {
		list.PushFront(val)

		return true
	}

	if idx == list.len {
		list.Push(val)

		return true
	}

	newNode := &node[T]{val: val}
	prevNode, _ := list.findByIndex(idx - 1)

	newNode.next = prevNode.next
	prevNode.next = newNode

	list.len++

	return true
}

func (list *List[T]) Set(val T, idx int) bool {
	node, ok := list.findByIndex(idx)

	if !ok {
		return false
	}

	node.val = val

	return true
}

func (list *List[T]) RemoveByIndex(idx int) bool {
	if idx < 0 || idx >= list.len {
		return false
	}

	if idx == 0 {
		list.PopFront()

		return true
	}

	if idx == list.len-1 {
		list.Pop()

		return true
	}

	prevNode, _ := list.findByIndex(idx - 1)

	prevNode.next = prevNode.next.next

	list.len--

	return true
}

func (list *List[T]) IndexOf(val T) int {
	var idx int
	current := list.head

	for current != nil {
		if current.val == val {
			return idx
		}

		current = current.next
		idx++
	}

	return -1
}

func (list *List[T]) Contains(val T) bool {
	return list.IndexOf(val) != -1
}

func (list *List[T]) Reverse() *List[T] {
	var prev *node[T]
	var next *node[T]

	current := list.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head, list.tail = list.tail, list.head

	return list
}
