package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "My Test",
		})
	})

	// r.GET("/blog", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "")
	// })
	r.Run() // listen and serve on 0.0.0.0:8080
}
