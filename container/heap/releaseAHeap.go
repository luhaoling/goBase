package heap

type MinHeap struct {
	Arr []int
}

func (c *MinHeap) Push(val int) {
	c.Arr = append(c.Arr, val)
	c.up(len(c.Arr) - 1)
}

func (c *MinHeap) Pop() (val int, ok bool) {
	if c.IsEmpty() {
		return 0, false
	}
	// 保存堆顶元素
	val = c.Arr[0]
	// 交换堆顶元素和末尾元素
	c.Arr[0], c.Arr[len(c.Arr)-1] = c.Arr[len(c.Arr)-1], c.Arr[0]
	// 移除末尾元素
	c.Arr = c.Arr[:len(c.Arr)-1]
	// 下沉调整
	c.down(0)
	return val, true
}

func (c *MinHeap) IsEmpty() bool {
	return len(c.Arr) == 0
}

// 重新调整堆的位置，把新插入的值放在堆中合适的位置
// 如果 i 位置的元素小于其父节点，那么将其替换
func (c *MinHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if c.Arr[i] < c.Arr[parent] {
			c.Arr[i], c.Arr[parent] = c.Arr[parent], c.Arr[i]
			i = parent
		} else {
			break
		}
	}
}

// 这一步是重新调整堆，即把末尾元素放在合适的位置，并且把最小的值移动到堆顶
func (c *MinHeap) down(i int) {
	n := len(c.Arr)
	for {
		left := 2*i + 1
		right := 2*i + 2
		minIdx := i

		// 检查左子节点是否小于当前值
		if left < n && c.Arr[left] < c.Arr[minIdx] {
			minIdx = left
		}

		// // 检查右子节点是否小于当前值
		if right < n && c.Arr[right] < c.Arr[minIdx] {
			minIdx = right
		}
		// 当前节点已经是最小值，说明已满足堆的性质
		if minIdx == i {
			break
		}
		// 将较小的子节点提升到父节点的位置，同时当前节点下沉到子节点位置
		c.Arr[i], c.Arr[minIdx] = c.Arr[minIdx], c.Arr[i]
		i = minIdx
	}
}

func (c *MinHeap) Init() {
	for i := len(c.Arr)/2 - 1; i >= 0; i-- {
		c.down(i)
	}
}

// 与 container/heap 包中堆内容的差别
// releaseAHeap 实现的堆的特点
// 1. 直接针对 int 类型的最小堆实现，无法拓展其他数据类型
// 2. 通过 Push() 和 Pop() 直接操作切片，方法参数和返回值类型固定
// 2. 通过 up() 和 down() 的固定逻辑强制维护最小堆，无法通过修改比较逻辑实现最大堆或其他堆类型
// 标准库的设计
// 1. 通过 heap.Interface{} 接口抽象堆操，支持任意类型和堆类型（如最小堆，最大堆，优先队列）
// 2. 通过接口方法调用实现堆操作，用户需自行定义 Less 方法控制堆属性
// 3. 不限制底层存储结构，用户可以用切片、链表或其他方法实现堆
