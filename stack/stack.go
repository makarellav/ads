package stack

type Stack[T comparable] struct {
	head *node[T]
	len  int
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) *Stack[T] {
	newNode := &node[T]{val: val}

	s.len++

	if s.head == nil {
		s.head = newNode

		return s
	}

	newNode.prev = s.head
	s.head = newNode

	return s
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if s.head == nil {
		var zero T

		return zero, false
	}

	oldHead := s.head
	s.head = s.head.prev

	s.len--

	return oldHead.val, true
}

func (s *Stack[T]) Peek() (v T, ok bool) {
	if s.head == nil {
		var zero T

		return zero, false
	}

	return s.head.val, true
}

func (s *Stack[T]) Len() int {
	return s.len
}
