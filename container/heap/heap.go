package heap

import (
	"container/heap"
	"sort"
)

// 数据流中第 k 大的元素
// 实现了一个优先队列进行解题

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

// 滑动窗口最大值
// 维护一个大顶堆进行解题
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(h.IntSlice)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// todo(Done) 理解 heap 包的实现原理以及它的设计原理。具体指的是它为什么能够让下述内容可以灵活调用 heap 包中的内容

// 操作分析
// Push 先调用用户实现的 Push(x) 添加元素，再调用内部的 up() 维护堆的性质。分离数据操作与堆调整逻辑，支持拓展性
// Pop 调用用户实现的 Swap()，再调用 down()，最后调用用户 Pop() 移除末尾。通过接口方法确保堆调整与数据操作的解耦

// 如何使用
// 用户需要自行实现 Push,Pop 方法所对应的数据操作

// 这种设计的优势
// 1. 接口抽象与拓展性
// 解耦数据与操作，将堆的存储结构、比较逻辑、增删操作通过接口分离。用户可以自由组合实现不同堆类型
// 支持复杂场景：可通过 Less 方法定义优先级，自定义元素类型等
// 2. 避免硬编码堆性质
// 通过 Less 方法的实现动态控制堆类型（最小堆或最大堆），而非在代码中固定比较方向
// 3. 性能与安全性的平衡
// 类型安全：用户需要自行处理类型断言，避免了标准库强制类型限制带来的不灵活
// 错误处理：用户需自行处理空堆等边界条件
