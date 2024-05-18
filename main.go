package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ronniesong0809/tinyKv/handler"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Root)
	router.GET("/kv/:key", handlers.GetValue)
	router.POST("/kv/:key", handlers.SetValue)
	router.PUT("/kv/:key", handlers.UpdateValue)
	router.DELETE("/kv/:key", handlers.DeleteValue)

	router.Run(":3000")
}
