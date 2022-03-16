package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pawnjiang/go-web/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)

	r.POST("/api/user/register", controller.Register)
	r.Run()
}
