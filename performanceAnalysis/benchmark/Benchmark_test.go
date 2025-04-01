package benchmark

import (
	"testing"
	"unsafe"
)

// -bench 表示需要benchmark运行的方法,.表示运行本目录所有Benchmark开头的方法
// -benchmem 显示与内存分配相关的详细信息
// -benchtime 设定每个基准测试用例的运行时间
// -cpuprofile 生成 CPU 性能分析文件
// -memprofile 生成内存性能分析文件
//> go test -bench='.' -benchmem -benchtime=10s -cpuprofile='cpu.prof' -memprofile='mem.prof'

func BenchmarkBytes2StrRaw(b *testing.B) {
	aa := []byte("abcdefg")
	for n := 0; n < b.N; n++ {
		Bytes2StrRaw(aa)
	}
}

func BenchmarkBytes2StrUnsafe(b *testing.B) {
	aa := []byte("abcdefg")
	for n := 0; n < b.N; n++ {
		Bytes2StrUnsafe(aa)
	}
}

func Bytes2StrRaw(b []byte) string {
	return string(b)
}

// 经验证，效率更高
func Bytes2StrUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
