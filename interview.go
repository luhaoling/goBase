package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	a := []string{"a", "b", "c"}
	// +
	ret := a[0] + a[1] + a[2]
	// fmt.Sprintf()
	ret = fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
	fmt.Println(ret)

	// strings.Builder
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	fmt.Println(sb.String())

	// bytes.Buffer
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	fmt.Println(buf.String())

	ret = strings.Join(a, "")

	type Set map[string]struct{}
	set := make(Set)
	for _, item := range []string{"A", "B", "C"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set))
	if _, ok := set["A"]; ok {
		fmt.Println("ok")
	}

}
