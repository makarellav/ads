package queue

import "github.com/makarellav/ads/lists/sll"

type Queue[T comparable] struct {
	list *sll.List[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{list: sll.New[T]()}
}

func (q *Queue[T]) Enqueue(val T) *Queue[T] {
	q.list.Push(val)

	return q
}

func (q *Queue[T]) Deque() (v T, ok bool) {
	return q.list.PopFront()
}

func (q *Queue[T]) Peek() (v T, ok bool) {
	return q.list.FindByIndex(0)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}
