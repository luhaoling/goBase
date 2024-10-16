package main

import (
	"fmt"
	"reflect"
)

type animal struct {
}

func (a *animal) Print(name string) {
	fmt.Println(name)
}

func main() {
	a := &animal{}
	t := reflect.ValueOf(a)
	method := t.MethodByName("Print")
	method.Call([]reflect.Value{reflect.ValueOf("hello")})
}
