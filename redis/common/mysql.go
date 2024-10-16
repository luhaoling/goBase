package common

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `column:"id"`
	UserName string `column:"user_name"`
	RealName string `column:"real_name"`
	Sex      string `column:"user_sex"`
	Phone    string `column:"user_phone"`
	Intro    string `column:"user_text"`
	UserType string `column:"user_type"`
}

type DatabaseInterface interface {
	Insert(ctx context.Context, user *User) error
	Get(ctx context.Context, id int64) (*User, error)
}

type Database struct {
	m *gorm.DB
}

func (d Database) Insert(ctx context.Context, user *User) error {
	return d.m.Create(user).Error
}

func (d Database) Get(ctx context.Context, id int64) (*User, error) {
	var user *User
	return user, d.m.First(&user, id).Error
}

func NewDatabase() DatabaseInterface {
	db, err := gorm.Open(mysql.Open("root:20010214@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Database{m: db}
}

func (u *User) TableName() string {
	return "t_user"
}
