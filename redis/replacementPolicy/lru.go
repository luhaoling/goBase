package replacementPolicy

import "fmt"

// 最近最少使用算法，根据最近一段时间内最长时间未使用的特点来淘汰对象
type LRU interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	GetArr()
}

func NewNode(key, val string) *LRUNode {
	return &LRUNode{
		key:  key,
		val:  val,
		next: nil,
		prev: nil,
	}
}

type LinkedList struct {
	LRULinkedList *LRUNode
	len           int
	cap           int
}

func NewLinkedList(cap int) *LinkedList {
	return &LinkedList{
		LRULinkedList: nil,
		len:           0,
		cap:           cap,
	}
}

type LRUNode struct {
	key  string
	val  string
	next *LRUNode
	prev *LRUNode
}

func (L *LinkedList) Set(key string, value string) error {
	lru := NewNode(key, value)
	last := FindLast(L.LRULinkedList)
	if L.len <= L.cap {
		last.next = lru
		lru.prev = last
		L.len++
		return nil
	}
	lru.prev = last.prev
	last.prev.next = lru
	return nil
}

func (L *LinkedList) Get(key string) (string, error) {
	head := L.LRULinkedList
	linked := L.LRULinkedList
	for linked != nil {
		if linked.key == key {
			linked.prev = linked.next
			linked.next.prev = linked.prev
			temp := head.next
			temp.prev = linked
			head.next = linked
			linked.next = temp
			return linked.val, nil
		}
		linked = linked.next
	}
	return "", NOTEXIT
}

func (L *LinkedList) GetArr() {
	linked := L.LRULinkedList
	for linked != nil {
		fmt.Println(linked.key, linked.val)
		linked = linked.next
	}
}

func FindLast(lruNode *LRUNode) *LRUNode {
	linked := lruNode
	for linked.next != nil {
		linked = linked.next
	}
	return linked
}
