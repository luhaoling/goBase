package replacementPolicy

import "container/list"

type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	c.list.MoveToFront(node)
	return node.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if node := c.keyToNode[key]; node != nil {
		node.Value = entry{key, value}
		c.list.MoveToFront(node)
		return
	}

	c.keyToNode[key] = c.list.PushFront(entry{key, value})
	if len(c.keyToNode) > c.capacity {
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

