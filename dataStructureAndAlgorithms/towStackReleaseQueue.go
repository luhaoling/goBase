package dataStructureAndAlgorithms

import "fmt"

type MyQueue struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		queue1: []int{},
		queue2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.queue2 = append(this.queue2, x)
}

func (this *MyQueue) Pop() int {
	if len(this.queue1) != 0 {
		val := this.queue1[0]
		this.queue1 = this.queue1[1:]
		return val
	}
	for _, v := range this.queue2 {
		this.queue1 = append(this.queue1, v)
	}
	this.queue2 = []int{}
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.queue1) != 0 {
		val := this.queue1[0]
		return val
	}
	for _, v := range this.queue2 {
		this.queue1 = append(this.queue1, v)
	}
	this.queue2 = []int{}
	val := this.queue1[0]
	return val
}

func (this *MyQueue) Empty() bool {
	fmt.Println(this.queue1)
	fmt.Println(this.queue2)
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
