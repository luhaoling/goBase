package linkList

import (
	"container/list"
	"fmt"
)

func LinkList() {
	link := list.New()

	for i := 0; i < 10; i++ {
		link.PushBack(i)
	}

	for p := link.Front(); p != nil; p = p.Next() {
		fmt.Println("number", p.Value)
	}
	for p := link.Back(); p != nil; p = p.Prev() {
		fmt.Println("number", p.Value)
	}
}
