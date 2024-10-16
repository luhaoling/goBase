package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {
	var ipArray [4]byte
	// 假设我们有一个 IPv4 地址 "192.168.1.1"
	ip := net.ParseIP("192.168.1.1").To4() // 转换为 IPv4 字节切片
	if ip == nil {
		fmt.Println("Not an IPv4 address")
	} else {
		copy(ipArray[:], ip) // 将切片复制到数组中
		fmt.Printf("IPv4 Address in Array: %d.%d.%d.%d\n", ipArray[0], ipArray[1], ipArray[2], ipArray[3])
	}
	fmt.Println(reflect.TypeOf(ipArray).Kind())
	fmt.Println(reflect.TypeOf(ipArray[:]).Kind())

}