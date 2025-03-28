package main

import "fmt"

func main() {
	f := getFruit("apple")
	fmt.Println(f.Fruit())
}

type FruitFactory interface {
	Fruit() string
}

func getFruit(t string) FruitFactory {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}

	}
	return nil
}

type apple struct{}

func (*apple) Fruit() string {
	return "吃苹果"
}

type banana struct{}

func (*banana) Fruit() string {
	return "吃香蕉"
}
