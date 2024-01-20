package sll

type List[T comparable] struct {
	head *node[T]
	tail *node[T]
	Len  int
}

func (lst *List[T]) Push(val T) *List[T] {
	newNode := &node[T]{val: val}

	if lst.head == nil {
		lst.head = newNode
		lst.tail = lst.head
	} else {
		lst.tail.next = newNode
		lst.tail = newNode
	}

	lst.Len++

	return lst
}

func (lst *List[T]) PopFront() (v T, ok bool) {
	if lst.head == nil {
		return zeroValue[T](), false
	}

	oldHead := lst.head

	lst.head = oldHead.next

	lst.Len--

	if lst.Len == 0 {
		lst.tail = nil
	}

	return oldHead.val, true
}

func (lst *List[T]) Pop() (v T, ok bool) {
	if lst.tail == nil {
		return zeroValue[T](), false
	}

	current := lst.head
	newTail := current

	for current.next != nil {
		newTail = current
		current = current.next
	}

	lst.tail = newTail
	lst.tail.next = nil

	lst.Len--

	if lst.Len == 0 {
		lst.head = nil
		lst.tail = nil
	}

	return current.val, true
}

func (lst *List[T]) PushFront(val T) *List[T] {
	newNode := &node[T]{val: val}

	if lst.Len == 0 {
		lst.head = newNode
		lst.tail = lst.head
	} else {
		newNode.next = lst.head
		lst.head = newNode
	}

	lst.Len++

	return lst
}

func (lst *List[T]) findByIndex(idx int) (n *node[T], ok bool) {
	if idx < 0 || idx >= lst.Len {
		return &node[T]{}, false
	}

	current := lst.head

	for i := 0; i < idx; i++ {
		current = current.next
	}

	return current, true
}

func (lst *List[T]) FindByIndex(idx int) (v T, ok bool) {
	n, found := lst.findByIndex(idx)

	if !found {
		return zeroValue[T](), false
	}

	return n.val, true
}

func (lst *List[T]) Insert(val T, idx int) bool {
	if idx < 0 || idx > lst.Len {
		return false
	}

	if idx == 0 {
		lst.PushFront(val)

		return true
	}

	if idx == lst.Len {
		lst.Push(val)

		return true
	}

	newNode := &node[T]{val: val}
	prevNode, _ := lst.findByIndex(idx - 1)

	newNode.next = prevNode.next
	prevNode.next = newNode

	lst.Len++

	return true
}

func (lst *List[T]) Set(val T, idx int) bool {
	node, ok := lst.findByIndex(idx)

	if !ok {
		return false
	}

	node.val = val

	return true
}

func (lst *List[T]) RemoveByIndex(idx int) bool {
	if idx < 0 || idx >= lst.Len {
		return false
	}

	if idx == 0 {
		lst.PopFront()

		return true
	}

	if idx == lst.Len-1 {
		lst.Pop()

		return true
	}

	prevNode, _ := lst.findByIndex(idx - 1)

	prevNode.next = prevNode.next.next

	lst.Len--

	return true
}

func (lst *List[T]) IndexOf(val T) int {
	var idx int
	current := lst.head

	for current != nil {
		if current.val == val {
			return idx
		}

		current = current.next
		idx++
	}

	return -1
}

func (lst *List[T]) Contains(val T) bool {
	return lst.IndexOf(val) != -1
}

func (lst *List[T]) Reverse() *List[T] {
	var prev *node[T] = nil
	var next *node[T] = nil

	current := lst.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	lst.head, lst.tail = lst.tail, lst.head

	return lst
}

func New[T comparable]() *List[T] {
	return &List[T]{}
}
