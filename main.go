package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)




func main() {
	r:=gin.Default()
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})
	r.POST("/login", func(c *gin.Context) {
		db:=connect()
		username := c.PostForm("username")
		password := c.PostForm("password")
		var uu = UserInfo{}
		result:=db.Find(&uu, "username=? AND password=?",username,password)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.HTML(200, "login_failed.html", nil)
		}else{
			c.HTML(200,"login_success.html", nil)
		}
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200,"register.html",nil)
	})
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		var uu = UserInfo{}
		db:=connect()
		result:=db.Find(&uu, "username=? AND password=?",username,password)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			userinfo :=UserInfo{username,password}
			db.Create(&userinfo)
			c.HTML(200,"register_success.html",nil)
		}else{
			c.HTML(200,"register_failed.html",nil)
		}
		defer db.Close()
	})
	r.Run(":8080")
}