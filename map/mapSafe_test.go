package main

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

type Elem1 struct {
	Key string
	Val int
}

func BenchmarkMapSafe(b *testing.B) {
	num := 10000
	testCase := genNoRepeatTestCase1(num)
	m := NewMapRMutex()
	for _, v := range testCase {
		m.Set(v.Key, v.Val)
	}
	b.ResetTimer()

	for i := 0; i < 5; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.N = 1000000

			wg := sync.WaitGroup{}
			wg.Add(b.N * 2)
			for i := 0; i < b.N; i++ {
				e := testCase[rand.Intn(num)]

				go func(key string, val interface{}) {
					m.Set(key, val)
					wg.Done()
				}(e.Key, e.Val)

				go func(key string) {
					_, _ = m.Get(key)
					wg.Done()
				}(e.Key)
			}
			wg.Wait()
		})
	}
}

func genNoRepeatTestCase1(num int) []Elem1 {
	elem := make([]Elem1, num)
	for i := 0; i < num; i++ {
		elem[i] = Elem1{
			Key: strconv.Itoa(i),
			Val: i,
		}
	}
	return elem
}
