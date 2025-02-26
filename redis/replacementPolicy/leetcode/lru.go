package leetcode

/*
double linked list release the lru algorithms
*/

type LRUCache struct{
	size int
	capacity int
	cache map[int]*DlinkedNode
	head,tail *DlinkedNode
}

type DlinkedNode struct{
	key,value int
	prev,next *DlinkedNode
}

func initDlinkedNode(key,value int)*DlinkedNode{
	return &DlinkedNode{
		key:key,
		value:value,
	}
}

func Constructor(capacity int)LRUCache{
	l:=LRUCache{
		cache:map[int]*DlinkedNode{},
		head:initDlinkedNode(0,0),
		tail:initDlinkedNode(0,0),
		capacity:capacity,
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	return l
}

func (this *LRUCache)Get(key int)int{
	if _,ok:=this.cache[key];!ok{
		return -1
	}
	node:=this.cache[key]
	this.moveTohead(node)
	return node.value
}

func (this *LRUCache)Put(key int,value int){
	if _,ok:=this.cache[key];!ok{
		node:=initDlinkedNode(key,value)
		this.cache[key]=node
		this.addToHead(node)
		this.size++
		if this.size>this.capacity{
			removed:=this.removeTail()
			delete(this.cache,removed.key)
			this.size--
		}
	}else{
		node:=this.cache[key]
		node.value=value
		this.moveTohead(node)
	}
}

func (this *LRUCache)addToHead(node *DlinkedNode){
	node.prev=this.head
	node.next=this.head.next
	this.head.next.prev=node
	this.head.next=node
}

func (this *LRUCache)removeNode(node *DlinkedNode){
	node.prev.next=node.next
	node.next.prev=node.prev
}

func (this *LRUCache)moveTohead(node *DlinkedNode){
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache)removeTail()*DlinkedNode{
	node:=this.tail.prev
	this.removeNode(node)
	return node
}


