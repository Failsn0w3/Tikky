package main

import (
	"Tikky-main-2/controller"
	"Tikky-main-2/database"
	"Tikky-main-2/mid"
	"github.com/gin-gonic/gin"
)

func main() {
	//token
	x := gin.Default()
	x.POST("/auth", controller.AuthHandler)
	x.GET("/home", mid.JWTAuthMiddleware(), homeHandler)
	x.Run(":9000")
	//token
	//获取初始化的数据库
	db := database.InitDB()
	//延迟关闭数据库
	defer db.Close()
	r := gin.Default()
	//注册

	//在9090端口启动服务
	panic(r.Run(":9090"))

}
