package sll

import (
	"fmt"
	"testing"
)

func assertLength(t testing.TB, want int, got int) {
	t.Helper()

	if got != want {
		t.Errorf("want length %d, got %d", want, got)
	}
}

func assertCorrectValue[T comparable](t testing.TB, nodeName string, want T, got T) {
	t.Helper()

	if got != want {
		t.Errorf("want %s value to be %v, got %v", nodeName, want, got)
	}
}

func assertLastElement[T comparable](t testing.TB, tail *node[T]) {
	t.Helper()

	if tail.next != nil {
		t.Errorf("want tail to have no next, got %v", tail.next.val)
	}
}

func TestList_Push(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		newVal := 1
		list := List[int]{}

		list.Push(newVal)

		wantLen := 1

		assertLength(t, wantLen, list.Len())
		assertCorrectValue(t, "tail", newVal, list.tail.val)
		assertLastElement(t, list.tail)

		if list.head != list.tail {
			t.Errorf("want head == tail, got %t", list.head == list.tail)
		}
	})

	t.Run("push multiple elements", func(t *testing.T) {
		list := List[string]{}
		newVals := []string{"a", "b", "c", "d"}

		for _, val := range newVals {
			list.Push(val)
		}

		assertLength(t, len(newVals), list.Len())
		assertCorrectValue(t, "tail", newVals[len(newVals)-1], list.tail.val)
		assertCorrectValue(t, "head", newVals[0], list.head.val)
		assertLastElement(t, list.tail)

		current := list.head
		for i := 0; i < len(newVals)-1; i++ {
			if current.val != newVals[i] {
				assertCorrectValue(t, fmt.Sprintf("list[%d]", i), newVals[i], current.val)
			}

			if current.next.val != newVals[i+1] {
				assertCorrectValue(t, fmt.Sprintf("list[%d].next", i), newVals[i+1], current.next.val)
			}

			current = current.next
		}
	})
}
