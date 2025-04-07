package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// go test -bench='.' -benchmem -benchtime=10s -cpuprofile='cpu.prof' -memprofile='mem.prof'
var slices []int

func TestSize(t *testing.T) {
	fmt.Printf("struct {} size:%d byte\n", unsafe.Sizeof(struct{}{}))
}

func init() {
	slices = getRawSlices()
}

func Benchmark1GetRawSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getRawSet(slices)
	}
}

func Benchmark1GetEmptyStructSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmptyStructSet(slices)
	}
}
