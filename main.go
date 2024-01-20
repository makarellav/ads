package main

import (
	"fmt"
	"github.com/makarellav/ads/lists/sll"
)

func main() {
	singlyLinkedList := sll.New[int]()

	singlyLinkedList.Push(10).Push(321)

	fmt.Println(singlyLinkedList.IndexOf(321))
}
