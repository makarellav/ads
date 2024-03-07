package stack

type node[T comparable] struct {
	val  T
	prev *node[T]
}
