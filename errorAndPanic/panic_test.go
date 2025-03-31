package errorAndPanic

import (
	"fmt"
	"testing"
)

// recover 恢复所在 goroutine 的 panic(恢复后，并不会接着继续从引发 panic 的地方继续执行，而是执行 defer 后面的函数)
func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Println("defer func")
	}()
	panic("This is panic")

	fmt.Println("nihao")
}

// recover 不能恢复其他 goroutine 的 panic
func TestOtherPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v", err)
		}
		return
	}()
	go func() {
		panic("panic")
	}()

	fmt.Println("nihao")
}
