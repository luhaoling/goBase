package main

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

type Elem struct {
	Key string
	Val int
}

func BenchmarkMapShared(b *testing.B) {
	num := 10000
	testCase := genNoRepeatTestCase(num)
	m := New()
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

func genNoRepeatTestCase(num int) []Elem {
	elem := make([]Elem, num)
	for i := 0; i < num; i++ {
		elem[i] = Elem{
			Key: strconv.Itoa(i),
			Val: i,
		}
	}
	return elem
}
