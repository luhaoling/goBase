package model

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Age      int64  `json:"age"`
}
