package heap

import "fmt"

func SetRelease() {
	set := make(map[string]struct{})
	set["string"] = struct{}{}
	set["string"] = struct{}{}
	set["str"] = struct{}{}
	set["Foo"] = struct{}{}
	for v := range set {
		fmt.Println(v)
	}
	delete(set, "Foo")
	fmt.Println(len(set))
	_, ok := set["Foo"]
	fmt.Println(ok)
}

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
)

// 设置
func Set(b, flag Bits) Bits {
	return b | flag
}

// 清除
func Clear(b, flag Bits) Bits {
	return b &^ flag
}

// 切换
func Toggle(b, flag Bits) Bits {
	return b ^ flag
}

// 检查
func Has(b, flag Bits) bool {
	return b&flag != 0
}

func BitSetRelease() {
	var b Bits
	b = Set(b, F0)
	b = Toggle(b, F2)
	for i, flag := range []Bits{F0, F1, F2} {
		fmt.Println(i, Has(b, flag))
	}
}
