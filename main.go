package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/image/:text", GetImageByText)
	r.Run(":8080")
}
