package linkList

import "fmt"

// 双向链表的定义：链表节点之间有前驱和后继两个指针，特别的是头节点的前驱指针为空，尾节点的后继指针为空

type DoubleListNode struct {
	Val  int
	Prev *DoubleListNode // 前驱
	Next *DoubleListNode // 后继
}

type DoubleLinkList struct {
	Head *DoubleListNode // 链表头指针
	Tail *DoubleListNode // 链表尾指针
}

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{Head: nil, Tail: nil}
}

// insertNode 既可以在头部插入，也可以在尾部插入
func (l *DoubleLinkList) insertNode(val int, atHead bool) {
	newNode := &DoubleListNode{Val: val}

	// 如果链表为空，新节点就是头部也是尾部
	if l.Head == nil {
		newNode.Prev = nil // 新节点没有前驱
		newNode.Next = nil // 新节点没有后继
		l.Head = newNode   // 头部指向新节点
		l.Tail = newNode   // 尾部指向新节点
		return
	}

	// 头插法
	if atHead {
		newNode.Next = l.Head
		newNode.Prev = nil
		l.Head.Prev = newNode
		l.Head = newNode
		// 尾插法
	} else {
		newNode.Prev = l.Tail
		newNode.Next = nil
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

// InsertAtHead 使用头插法在链表头部插入节点
func (l *DoubleLinkList) InsertAtHead(val int) {
	l.insertNode(val, true)
}

// InsertAtTail 使用尾插法在链表尾部插入节点
func (l *DoubleLinkList) InsertAtTail(val int) {
	l.insertNode(val, false)
}

// InsertAtRandomPosition 在随机位置插入节点
// 如果位置超出链表长度，则在尾部插入
func (l *DoubleLinkList) InsertAtRandomPosition(pos int, val int) {
	// 位置小于0，或者链表为空，直接在头部插入
	if pos <= 0 || l.Head == nil {
		l.InsertAtHead(val)
		return
	}

	// 计算链表长度
	length := 0
	current := l.Head
	for current != nil {
		length++
		current = current.Next
	}

	// 如果超出链表长度，在尾部插入
	if pos >= length {
		l.InsertAtTail(val)
		return
	}

	// 遍历链表，找到指定位置的前一个节点
	current = l.Head
	for i := 0; i < pos-1; i++ {
		current = current.Next
	}
	// 插入节点
	newNode := &DoubleListNode{Val: val}
	newNode.Next = current.Next
	newNode.Prev = current
	current.Next.Prev = newNode
	current.Next = newNode
}

func (l *DoubleLinkList) DeleteAtPosition(pos int) {
	if l.Head == nil || pos < 0 {
		return
	}

	if pos == 0 {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}

	current := l.Head
	for i := 0; current != nil && i < pos; i++ {
		current = current.Next
	}

	if current == nil {
		return
	}

	if current.Next == nil {
		current.Prev.Next = nil
	} else {
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}

	// 清理当前节点的指针，帮助垃圾回收
	current.Next = nil
	current.Prev = nil
}

// traverseList 遍历链表的函数
// 根据 reverse 的值决定遍历方向，并执行给定的操作 action
// reverse 为 true 时，从尾到头遍历；为 false 时，从头到尾部遍历
func (l *DoubleLinkList) traverseList(reverse bool, action func(node *DoubleListNode)) {
	if l.Head == nil {
		fmt.Println("链表为空")
		return
	}

	current := l.Head
	if reverse {
		current = l.Tail
		for {
			action(current)
			current = current.Prev
			if current == nil {
				break
			}
		}
	} else {
		for {
			action(current)
			current = current.Next
			if current == nil {
				break
			}
		}
	}
}

// PrintListForward 打印双向链表元素（从前向后）
func (l *DoubleLinkList) PrintListForward() {
	l.traverseList(false, func(node *DoubleListNode) {
		fmt.Println(node.Val)
	})
}

// PrintListBackward 打印双向链表元素（从后向前）
func (l *DoubleLinkList) PrintListBackward() {
	l.traverseList(true, func(node *DoubleListNode) {
		fmt.Println(node.Val)
	})
}

// 总结：
// 进行链表操作时要注意链表头部，链表尾部，空链表这几种情况
