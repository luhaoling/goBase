package dataStructureAndAlgorithms

// release a int array
// 要求，
// 插入操作：往数组中的任意地方插入一个数据
// 删除操作：可以删除数组中的任意一个元素
// 查找操作：可以根据数组下标返回数组中的任意一个元素

type ArrayInterface interface {
	Lookup(index int) int
	Insert(index ...int) bool
	update(index, value int) bool
	delete(index int) bool
}

type Array struct {
	len int
	arr []int
}

func (a *Array) Lookup(index int) int {
	if a.len >= index {
		return -1
	}
	return a.arr[index]
}

func (a *Array) Insert(index ...int) bool {
	if len(index) == 1 {
		a.arr = append(a.arr, index[0])
		return true
	} else if len(index) == 2 {
		ind := index[0]
		val := index[1]
		for i := a.len; i >= ind; i-- {
			a.arr[i] = a.arr[i-1]
		}
		a.arr[ind] = val
		return true
	}
	return false
}

func (a *Array) update(index, value int) bool {
	if a.len >= index {
		return false
	}
	a.arr[index] = value
	return true
}

func (a *Array) delete(index int) bool {
	if index > a.len {
		return false
	}
	for i := index; i < a.len; i++ {
		a.arr[i] = a.arr[i+1]
	}
	return true
}

func NewArray(cap int) ArrayInterface {
	return &Array{
		len: 0,
		arr: make([]int, cap),
	}
}
