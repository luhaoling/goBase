package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	age  int
}

func (u *User) GetName() string {
	return u.Name
}
func main() {
	//创建切片
	users := make([]*User, 1, 3)
	users[0] = &User{
		Name: "mszlu",
	}
	//修改切片 必须使用指针
	userSliceValue := reflect.ValueOf(&users)
	//我们可以改变切片的长度
	//userSliceValue.Elem().Index(1).Set(reflect.ValueOf(&User{Name: "mszlu"})) //会报错 slice index out of range
	userSliceValue.Elem().SetLen(2)
	userSliceValue.Elem().Index(1).Set(reflect.ValueOf(&User{Name: "mszlu1"}))
	fmt.Println(users[1].Name)

	//也可以直接Append
	userSliceValue = reflect.Append(userSliceValue.Elem(), reflect.ValueOf(&User{Name: "mszlu2"}))
	users = userSliceValue.Interface().([]*User)
	fmt.Println(users[2].Name)
}
