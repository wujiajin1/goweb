package main

import (
	"github.com/jinzhu/gorm"
)

// UserInfo 用户信息
type UserInfo struct {
	Username string `gorm:"primaryKey"`
	Password string
}

type BlogsInfo struct {
	User  string
	Title string
	Text  string `gorm:"type:text"`
	PS    string
	Time  string
}

func connect() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:WJJ99zyh@(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&BlogsInfo{})
	return db
}
