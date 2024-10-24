package replacementPolicy

import (
	"errors"
	"fmt"
	"sort"
)

var NOTEXIT error = errors.New("not exit")

// 最不经常使用算法，根据使用次数进行淘汰对象
type LFU interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	GetArr()
}

type arr struct {
	element []ele
	len     int
	cap     int
}

func (a *arr) GetArr() {
	fmt.Println(a.element)
}

func NewArr(cap int) LFU {
	return &arr{
		element: make([]ele, cap),
		len:     0,
		cap:     cap,
	}
}

type ele struct {
	key   string
	val   string
	count int
}

func (a *arr) Get(key string) (string, error) {
	for i := 0; i < a.len; i++ {
		if a.element[i].key == key {
			a.element[i].count++
			return a.element[i].val, nil
		}
	}
	return "", NOTEXIT
}

func (a *arr) Set(key string, val string) error {
	if a.len < a.cap {
		a.element[a.len].key = key
		a.element[a.len].val = val
		a.len++
		return nil
	}
	targetKey := replacement(a.element)
	for i := 0; i < a.len; i++ {
		if a.element[i].key == targetKey {
			a.element[i].key = key
			a.element[i].val = val
			return nil
		}
	}
	return nil
}

func replacement(arr []ele) string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})
	return arr[len(arr)-1].key
}
