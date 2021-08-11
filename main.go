package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

func main() {
	db:= connect()
	for i := 60; i <90 ; i++ {
		b:=BlogsInfo{
			Title: string(i)+string(i),
			Text:  string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i)+string(i),
			PS:    string(i)+string(i)+string(i)+string(i)+string(i),
		}
		db.Create(&b)
	}
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
				"usr":username,
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
			username,password="",""
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
		result := db.Find(&uu, "username=?", username)
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
				"usr":username,
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
		bloginfo := BlogsInfo{title, text, ps}
		db.Create(&bloginfo)
		c.HTML(200, "main.html", gin.H{
			"code": 0,
			"usr":  username,
		})
	})
	r.GET("/blogs", func(c *gin.Context) {
		p, _ := strconv.Atoi(c.Query("page"))
		I,BI:=pageDivision(p)
		c.HTML(200,"blogs.html",gin.H{
			"usr":username,
			"data":BI,
			"page":I,
		})
	})
	r.Run(":8080")
}
