package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.LoadHTMLFiles("index.html")

	r.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	_ = r.Run()
}
