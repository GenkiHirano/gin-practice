package main

import (
	"github.com/GenkiHirano/gin-practice/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/user", handler.CreateUser)

	r.Run()
}
