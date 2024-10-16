package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	slice := make([]int, 1, 10)
	slice2 := make([]int, 1, 9)
	fmt.Println(reflect.DeepEqual(slice, slice2))

	arr := []int{1, 2, 3}

	s := slices.Clone(arr)
	fmt.Println(s)
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p", s)
}
