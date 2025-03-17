package heap

import (
	"fmt"
	"testing"
)

func TestReleaseHeap(t *testing.T) {
	mh := MinHeap{}

	mh.Push(10)
	mh.Push(16)
	mh.Push(4)
	mh.Push(5)
	mh.Push(9)
	mh.Push(3)
	mh.Push(2)
	mh.Push(12)
	fmt.Println(mh.Arr)
	for !mh.IsEmpty() {
		fmt.Println(mh.Pop())
	}
}
