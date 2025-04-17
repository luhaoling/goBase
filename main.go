package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println(IsPalindrome("aba"))
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := GetUserInfo(id)
	if err != nil {
		c.JSON(500, &Resp{
			Code: "500",
			Msg:  "",
			Data: nil,
		})
	}
	c.JSON(http.StatusOK, &Resp{
		Code: "200",
		Msg:  "",
		Data: user,
	})
}

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserInfo(id int) (*User, error) {
	return &User{
		Id:   id,
		Name: "xiaoming",
	}, nil
}

func IsPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
