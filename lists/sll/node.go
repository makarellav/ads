package sll

type node[T comparable] struct {
	next *node[T]
	val  T
}
