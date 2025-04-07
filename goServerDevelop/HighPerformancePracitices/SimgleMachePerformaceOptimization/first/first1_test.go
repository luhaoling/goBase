package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// go test -bench='.' -benchmem -benchtime=10s -cpuprofile='cpu.prof' -memprofile='mem.prof'
var slices1 []int

func TestSize1(t *testing.T) {
	fmt.Printf("struct {} size:%d byte\n", unsafe.Sizeof(struct{}{}))
}

func init() {
	slices1 = getRawSlices()
}

func BenchmarkGetRawSet1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getRawSet(slices1)
	}
}

func BenchmarkGetEmptyStructSet1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getCapacitySet(slices1)
	}
}
