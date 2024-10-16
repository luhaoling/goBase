package main

import "fmt"

func main() {
	apple := appleFactory{}
	fmt.Println(apple.Fruit())

	banana := bananaFactory{}
	fmt.Println(banana.Fruit())
}

type Fruit1 interface {
	Fruit() string
}

type appleFactory struct{}

func (*appleFactory) Fruit() string {
	return "吃苹果"
}

type bananaFactory struct{}

func (*bananaFactory) Fruit() string {
	return "吃香蕉"
}
