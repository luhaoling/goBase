package dataStructureAndAlgorithms

import "sync"

/*
base int, release a concurrency safe queue.
*/
type Queue struct {
	sync.Mutex
	len int
	cap int
	arr []int
}

func (q *Queue) Poll() int {
	q.Lock()
	defer q.Unlock()
	if q.len > 0 {
		q.len--
		return q.arr[0]
	}
	return -1
}

func (q *Queue) Push(i int) int {
	q.Lock()
	defer q.Unlock()
	if q.cap > q.len {
		q.arr[q.len] = i
		q.len++
		return i
	}
	return -1
}

func (q *Queue) getCap() int {
	return q.cap
}

func (q *Queue) getLen() int {
	return q.len
}

type QueueInterface interface {
	Poll() int
	Push(int) int
	getCap() int
	getLen() int
}

func NewQueue(cap int) QueueInterface {
	return &Queue{
		cap: cap,
		arr: make([]int, 10),
	}
}
