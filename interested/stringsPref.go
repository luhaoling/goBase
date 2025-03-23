package interested

import "strings"

// judge a string is another string's pref
// 这也是库函数 hasPref 的实现方法
func JudgeStringPref(words []string, s string) int {
	ans := 0
	for _, word := range words {
		if len(s) <= len(word) && s[:len(word)] == word {
			ans++
		}
	}
	return ans
}

func JudgeStringPref1(words []string, s string) int {
	res := 0
	for _, word := range words {
		if strings.HasPrefix(word, s) {
			res++
		}
	}
	return res
}
