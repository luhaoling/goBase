package second

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"testing"
)

// go test -run=none -bench=BenchmarkGenerateIdsBuilder -benchtime=10s -gcflags=all=-l -cpuprofile cpu.prof
// go tool pprof -http=":8081" cpu.prof

var users []*User

func init() {
	for i := 0; i < 1000; i++ {
		users = append(users, &User{Id: i, Name: fmt.Sprintf("user%d", i)})
	}
}

func BenchmarkGenerateIdsRaw(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsRaw(users)
	}
}

// strings.Builder usage
func TestStringBuilder(t *testing.T) {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}

func BenchmarkGenerateIdsRaw2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsRaw(users)
	}
}

func BenchmarkGenerateIdsBuidler(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateIdsBuidler(users)
	}
}

func BenchmarkA(b *testing.B) {

}

func smallestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	str := [26]int{}
	for _, v := range s[:len(s)/2] {
		str[v-'a']++
	}

	n := len(s)
	ans := make([]byte, 0, n)

	for i, c := range str {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(i)}, c)...)
	}
	t := slices.Clone(ans)

	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}

	slices.Reverse(t)
	return string(append(ans, t...))
}
