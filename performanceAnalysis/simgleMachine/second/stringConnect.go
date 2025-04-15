package second

import (
	"fmt"
	"strings"
)

type User struct {
	Id   int
	Name string
}

func GenerateIdsRaw(users []*User) (string, string, []byte) {
	name := ""
	idStr := ""
	var nameByte []byte
	for index := range users {
		idStr = fmt.Sprint(users[index].Id)
		name = name + "," + users[index].Name
		nameByte = []byte(users[index].Name)
	}
	return idStr, name, nameByte
}

func GenerateIdsBuidler(users []*User) (string, string, []byte) {
	names := ""
	idStr := ""
	var nameByte []byte
	length := 0
	for index := range users {
		idStr = fmt.Sprint(users[index].Id)
		nameByte = []byte(users[index].Name)
		length += len(users[index].Name) + 1
	}
	var builder strings.Builder
	builder.Grow(length)
	for index := range users {
		builder.WriteString(",")
		builder.WriteString(users[index].Name)
	}
	return idStr, names, nameByte
}
