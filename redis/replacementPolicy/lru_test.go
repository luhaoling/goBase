package replacementPolicy

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLinkedList(3)
	lru.Set("1", "1")
	val, _ := lru.Get("1")
	fmt.Println("1", val)
}
