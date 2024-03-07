package main

import (
	"fmt"
	"github.com/makarellav/ads/stack"
)

func main() {
	s := stack.New[int]()

	s.Push(1)
	s.Push(321)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	_, ok := s.Peek()

	fmt.Println(ok)
}
