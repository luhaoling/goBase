package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	slices := getCapacitySlices()
	getCapacitySet(slices)
	// 设置响应头，这里设置Content-Type为text/plain，表示返回纯文本内容
	w.Header().Set("Content-Type", "text/plain")
	// 向客户端写入响应内容
	fmt.Fprintln(w, "hh")
}
func main() {
	// 注册路由，当客户端访问根路径"/"时，会调用Handler函数进行处理
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// getRawSlices 循环往切片append数据
func getRawSlices() []int {
	n := 10000
	slices := make([]int, 0)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

func getCapacitySlices() []int {
	n := 10000
	slices := make([]int, 0, 10000)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

// 构造集合
func getRawSet(slices []int) map[int]bool {
	set := make(map[int]bool, 0)
	for _, item := range slices {
		set[item] = true
	}
	return set
}

func getEmptyStructSet(slices []int) map[int]struct{} {
	set := make(map[int]struct{}, 0)
	for _, item := range slices {
		set[item] = struct{}{}
	}
	return set
}

func getCapacitySet(slices []int) map[int]struct{} {
	set := make(map[int]struct{}, len(slices))
	for _, item := range slices {
		set[item] = struct{}{}
	}
	return set
}

// ab 压测
// -n 请求数量
// -c 并发数量
// ab -n 30000 -c 2 http://127.0.0.1:8888/

// 内存火焰图
// go tool pprof -http :8889 http://127.0.0.1:8888/debug/pprof/heap

// 优化思路：
// 当使用 map 构造集合时，可以将 value 类型设置为空结构体，空结构体不占用内存空间。这样可以帮我们降低内存消耗
// 当创建 map 和切片对象时，如果可以提前确定容器容量，就可以传入 make 函数中，从而避免往集合中添加数据时出发扩容迁移，达到降低内存和 CPU 资源消耗的目的

// 优化 tips:
// 利用火焰图查看内存占用的时候，记得思考一下，当前代码的想法究竟对不对？
// 感觉代码思路没问题，再思考一下，当存储的数据量达到一定程度的时候，思考一下使用的数据类型的底层数据结构会不会对内存有较大的影响
