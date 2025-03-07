package dataStructureAndAlgorithms

// release a int linked
// 要求：
// 可以插入一个元素（包括插入一个指定位置的元素）
// 可以删除指定元素
// 可以进行遍历

type LinkedInterface interface {
	Insert(int, int) bool
	Delete(int) bool
	lookup(int) (int, bool)
}

type LinkedList struct {
	dummy *Node
	len   int
}

func (l *LinkedList) Insert(i int, i2 int) bool {
	if l.len < i {
		return false
	} else if l.len > i || i == 0 {
		head := l.dummy
		for ; head != nil; head = head.Next {
		}
		head.Next = &Node{
			val:  i2,
			Next: nil,
		}
	} else {
		head := l.dummy

		for in := 0; in < i; in++ {
			head = head.Next
		}
		nxt := head.Next
		head.Next = &Node{
			val:  i2,
			Next: nxt,
		}
	}
	return true
}

func (l *LinkedList) Delete(i int) bool {
	if l.len < i {
		return false
	}
	head := l.dummy
	for in := 0; in < i; in++ {
		head = head.Next
	}
	head.Next = head.Next.Next
	return true
}

func (l *LinkedList) lookup(i int) (int, bool) {
	if l.len < i {
		return 0, false
	}
	head := l.dummy
	for in := 0; in <= i; in++ {
		head = head.Next
	}
	return head.val, true
}

type Node struct {
	val  int
	Next *Node
}

func NewLinkedList() LinkedInterface {
	return &LinkedList{
		dummy: &Node{
			val:  0,
			Next: nil,
		},
		len: 0,
	}
}
