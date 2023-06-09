package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("/templates/*")

	router.GET("/", getData)
	router.Run(":8080")
}

func getData(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"image": "http://nginx:80/images/example.jpg",
	})
}
