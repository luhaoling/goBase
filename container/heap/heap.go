package heap

import (
	"container/heap"
	"sort"
)

// 数据流中第 k 大的元素
// 实现了一个优先队列进行解题
// todo 理解 heap 包的实现原理以及它的设计原理。具体指的是它为什么能够让下述内容可以灵活调用 heap 包中的内容

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

// 关于 Heap 包
// 它为任何实现 heap.Interface{} 的类型提供堆操作。默认支持最小堆，但是可以通过自定义 Less 方法实现最大堆。
