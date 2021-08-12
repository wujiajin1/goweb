package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"time"
)

func main() {
	//db:=connect()
	//for i := 0; i < 10; i++ {
	//
	//	for j := 65; j < 122; j++ {
	//		r:=string(j)
	//		for w:=i; w > 0 ; w-- {
	//			r+=string(j)
	//		}
	//		b:=BlogsInfo{
	//			User:  r,
	//			Title: r,
	//			Text:  r,
	//			PS:    r,
	//			Time:  r,
	//		}
	//		db.Create(&b)
	//	}
	//}
	var username, password string
	r := gin.Default()
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"code": 0,
		})
	})
	r.GET("/main", func(c *gin.Context) {
		if username != "" {
			c.HTML(200, "main.html", gin.H{
				"code": 1,
				"usr":  username,
			})
		} else {
			c.HTML(200, "main.html", gin.H{
				"code": 2,
			})
		}
	})
	r.POST("/main", func(c *gin.Context) {
		db := connect()
		username = c.PostForm("username")
		password = c.PostForm("password")
		var uu = UserInfo{}
		result := db.Find(&uu, "username=? AND password=?", username, password)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			username, password = "", ""
			c.HTML(200, "login.html", gin.H{
				"code": 1,
			})
		} else {
			c.HTML(200, "main.html", gin.H{
				"code": 2,
				"usr":  username,
			})
		}
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{
			"code": 0,
		})
	})
	r.POST("/register", func(c *gin.Context) {
		rigiName := c.PostForm("username")
		rigiPassword := c.PostForm("password")
		var uu = UserInfo{}
		db := connect()
		result := db.Find(&uu, "username=?", rigiName)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			userinfo := UserInfo{rigiName, rigiPassword}
			db.Create(&userinfo)
			c.HTML(200, "register.html", gin.H{
				"code": 2,
			})
		} else {
			c.HTML(200, "register.html", gin.H{
				"code": 1,
			})
		}
		defer db.Close()
	})
	r.GET("/write", func(c *gin.Context) {
		if username != "" {
			c.HTML(200, "writeblog.html", gin.H{
				"code": 1,
				"usr":  username,
			})
		} else {
			c.HTML(200, "writeblog.html", gin.H{
				"code": 2,
			})
		}
	})
	r.POST("/write", func(c *gin.Context) {
		title := c.PostForm("title")
		text := c.PostForm("text")
		ps := c.PostForm("ps")
		db := connect()
		timeUnix := time.Now().Unix() //已知的时间戳
		formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
		bloginfo := BlogsInfo{username, title, text, ps, formatTimeStr}
		db.Create(&bloginfo)
		c.HTML(200, "main.html", gin.H{
			"code": 0,
			"usr":  username,
		})
	})
	r.GET("/blogs", func(c *gin.Context) {
		p, _ := strconv.Atoi(c.Query("page"))
		I, BI := pageDivision(p)
		c.HTML(200, "blogs.html", gin.H{
			"usr":  username,
			"data": BI,
			"page": I,
		})
	})
	r.Run(":8080")
}
