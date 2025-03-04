package dataStructureAndAlgorithms

//
//import (
//	"math"
//	"math/rand"
//	"time"
//)
//
//const (
//	maxLevel    int     = 18
//	probability float64 = 1 / math.E
//)
//
//type handleEle func(e *Element) bool
//
//type (
//	Node struct {
//		next []*Element
//	}
//
//	Element struct {
//		Node
//		key   []byte
//		value interface{}
//	}
//
//	SkipList struct {
//		Node
//		maxLevel       int
//		Len            int
//		randSource     rand.Source
//		probability    float64
//		probTable      []float64
//		prevNodesCache []*Node
//	}
//)
//
//func NewSkipList() *SkipList {
//	return &SkipList{
//		Node:           Node{next: make([]*Element, maxLevel)},
//		prevNodesCache: make([]*Node, maxLevel),
//		maxLevel:       maxLevel,
//		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
//		probability:    probability,
//		probTable:      probabilityTable(probability, maxLevel),
//	}
//}
//
//func (e *Element) Key() []byte {
//	return e.key
//}
//
//func (e *Element) Value() interface{} {
//	return e.value
//}
//
//func (e *Element) SetValue(val interface{}) {
//	e.value = val
//}
//
//func (e *Element) Next() *Element {
//	return e.next[0]
//}
//
//func (t *SkipList) Front() *Element {
//	return t.next[0]
//}
