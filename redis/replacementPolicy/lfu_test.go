package replacementPolicy

import (
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {
	lfu := NewArr(3)
	lfu.Set("1", "1")
	val, _ := lfu.Get("1")
	fmt.Println("get \"1\":", val)
	lfu.GetArr()
	fmt.Println()
	lfu.Set("2", "2")
	fmt.Println("get \"2\":", val)
	lfu.Set("3", "3")
	lfu.Set("4", "4")
	lfu.GetArr()

}
