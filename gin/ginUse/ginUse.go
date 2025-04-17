package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", GetUser)
	r.Run(":8080")
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	user, err := getUserInfo(id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err,
			"data": nil,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  nil,
		"data": user,
	})
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	age  int64  `json:"age"`
}

func getUserInfo(id int) (*User, error) {
	return &User{
		Id:   id,
		Name: "name",
		age:  18,
	}, nil
}
