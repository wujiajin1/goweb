package main

import "github.com/jinzhu/gorm"

func connect()(db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:WJJ99zyh@(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}