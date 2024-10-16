package main

import (
	"fmt"
	"math"
	"sort"
)

var a func()

func main() {
	var x interface{} = nil
	var y *int = nil
	interfaceIsNil(x)
	interfaceIsNil(y)
	fmt.Println(x == y)
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func majorityElement(nums []int) int {
	mapKey := make(map[int]int)
	for _, v := range nums {
		nums, ok := mapKey[v]
		if ok {
			mapKey[v] = nums
		} else {
			mapKey[v] = 1
		}
	}
	sum := 0
	judge := len(nums) / 2
	for _, v := range mapKey {
		if judge >= v {
			sum += 1
		}
	}
	return sum
}

func deferTest() *int {
	i := 1
	defer func(x *int) {
		*x++
	}(&i)
	return &i
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func dfs(prices []int, i int, hold bool) int {
	if i < 0 {
		if hold {
			return math.MinInt
		}
		return 0
	}

	if hold {
		return max(dfs(prices, i-1, true), dfs(prices, i-1, false)-prices[i])
	}
	return max(dfs(prices, i-1, false), dfs(prices, i-1, true)+prices[i])

}

func reverse(nums []int, start, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
