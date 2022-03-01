package main

import (
	"github.com/GenkiHirano/gin-practice/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ginでユーザー登録のルーティングを書いてください。
	// パス/userがリクエストされた時に、ユーザー登録を処理してください。
	r.POST("/user", handler.CreateUser())

	// Ginでサーバー起動の処理を書いてください。
	r.Run()
}
