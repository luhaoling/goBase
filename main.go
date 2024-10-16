package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main start")

	go forloop()
	fmt.Println("wait")

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	fmt.Println("main over")

}

func forloop() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
