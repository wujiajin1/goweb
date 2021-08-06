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
		c.HTML(200,"login.html",gin.H{
			"code":0,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		db:=connect()
		username := c.PostForm("username")
		password := c.PostForm("password")
		var uu = UserInfo{}
		result:=db.Find(&uu, "username=? AND password=?",username,password)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.HTML(200, "login.html", gin.H{
				"code":1,
			})
		}else{
			c.HTML(200,"main.html", gin.H{
				"code":2,
				"usr":username,
			})
		}
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200,"register.html",gin.H{
			"code":0,
		})
	})
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		var uu = UserInfo{}
		db:=connect()
		result:=db.Find(&uu, "username=?",username)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			userinfo :=UserInfo{username,password}
			db.Create(&userinfo)
			c.HTML(200,"register.html",gin.H{
				"code":2,
			})
		}else{
			c.HTML(200,"register.html",gin.H{
				"code":1,
			})
		}
		defer db.Close()
	})
	r.GET("/write", func(c *gin.Context) {
		c.HTML(200,"writeblog.html",nil)
	})
	r.POST("write", func(c *gin.Context) {
		title  := c.PostForm("title")
		text := c.PostForm("text")
		ps := c.PostForm("ps")
		db:=connect()
		bloginfo:=BlogsInfo{title,text,ps}
		db.Create(&bloginfo)
		c.HTML(200,"main.html",nil)
	})
	r.Run(":8080")
}