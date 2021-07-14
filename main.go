package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// UserInfo 用户信息
type UserInfo struct {
	Username string
	Password string
}


func main() {
	r:=gin.Default()
	r.LoadHTMLFiles("./login.html","./register.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200,"login.html",nil)
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200,"register.html",nil)
	})
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("un")
		password := c.PostForm("pw")
		userinfo :=UserInfo{username,password}
		db:=connect()
		if db.Create(&userinfo).Error!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"user has been created",
			})
			return
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"ok",
		})
		defer db.Close()
	})
	r.GET("/index", func(c *gin.Context) {
	//	username := c.PostForm("un")
	//	password := c.PostForm("pw")
	//	userinfo :=UserInfo{username,password}
	//var uu UserInfo
	//	db:=connect()
	//	if db.Where("username=?", username).Find(&userinfo)!=nil{
	//	fmt.Println(db.Where(&UserInfo{Username: "p93002612", Password: "WJJ99zyh"}).Find(&uu).Error)
	//	}
	//	defer db.Close()
	})
	r.Run(":8080")
}