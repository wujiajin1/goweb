package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID uint
	Username string
	Password string
}


func main() {
	r:=gin.Default()
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200,"login.html",nil)
		//username := c.Query("un")
		//password := c.Query("pw")
		//db:=connect()

	})
	//defer db.Close()
	r.Run(":8080")
}