package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		log.Println("GET / called") // debugging proof
		c.HTML(200, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	r.Run(":8080")
}
