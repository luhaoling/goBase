package main

import (
	"fmt"
	"reflect"
)

type Slices[T int | float32 | float64 | uint | uint8 | string] []T

type Slice[T int | float32 | float64 | uint | uint8 | string] []T

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

type MyMap2[KEY int, VALUE float32] map[KEY]VALUE

type MyStruct[T int | string] struct {
	Name string
	Data T
}

type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

type MyChan[T int | string] chan T

type CommonType[T int | string | float32] struct {
}

type NewType[T interface{ *int }] []T

type NewType2[T interface{ *int | *float64 }] []T

type NewType3[T *int,] []T

type Wow[T int | string] int

type UintSlice[T uint | uint8] Slice[T]

type IntAndStringSlice[T int | string] Slice[T]

type IntSlice[T int] IntAndStringSlice[T]

type WowMap[T int | string] map[string]Slice[T]

type WowMap2[T Slice[int] | Slice[string]] map[string]T

type MySlice[T int | float32] []T

type Queue[T interface{}] struct {
	elements []T
}

func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

func (q Queue[T]) Size() int {
	return len(q.elements)
}

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

func (receiver Queue[T]) Put1(value T) {
	fmt.Println("%T", value)

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Int:
	//todo
	case reflect.String:
		// todo
	}
}

type IntUintFloat interface {
	int | uint | float32 | float64 | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

type SliceTe[T IntUintFloat] []T

func main() {
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("type name: %T", a)

	var b Slice[float32] = []float32{1., 2., 3.}
	fmt.Printf("type name: %T", b)

	fmt.Println("type a==b:")

	var c MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
	fmt.Println("MyMap is:", c)

	d := MyMap2[int, float32]{
		1: 9.6,
		2: 8.4,
	}
	fmt.Println("MyMap2 is:", d)

	var ws WowStruct[float32, []float32]
	var ws2 WowStruct[int, []int]
	fmt.Println(ws, ws2)

	var s MySlice[int] = []int{1, 2, 3, 4, 5}
	fmt.Println(s.Sum())

	var s2 MySlice[float32] = []float32{1., 2., 3., 4., 5.}
	fmt.Println(s2.Sum())

	Add(1, 2)
	Add[int, int](1, 2)
}

func MyFunc[T int | float32 | float64](a, b T) {
	fn2 := func(i, j T) T {
		return i*2 - j*2
	}
	fn2(a, b)
}

type A[T int | float32 | float64] struct {
}

func (receive A[T]) Add(a, b T) T {
	return a + b
}

type MyMap1[KEY comparable, VALUE any] map[KEY]VALUE
type Slice111[KEY comparable, VALUE Ordered] map[KEY]VALUE

type MyMap23[KEY comparable, VALUE any] map[KEY]VALUE
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | uintptr
}

type Float interface {
	~float32 | ~float64
}
type aaa[T Ordered] []T

type MySlice11[T int | float32 | float64] []T

func (s MySlice[T]) Sum11() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}
