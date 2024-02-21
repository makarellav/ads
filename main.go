package main

import (
	"fmt"
	"github.com/makarellav/ads/sort"
)

func main() {
	x := []int{6, 5, 4, 3, 2, 1}
	sort.BubbleSort(x)
	fmt.Println(x)
}
