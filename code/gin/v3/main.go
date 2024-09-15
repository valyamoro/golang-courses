package main 

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s!", name)
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		c.String(200, "Search query: %s", query)
	})
	r.Run()
}
