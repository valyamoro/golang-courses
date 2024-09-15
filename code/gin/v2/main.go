package main 

import "github.com/gin-gonic/gin"
import "fmt"

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		fmt.Printf("Status code: %d\n", status)
	}
}

func main() {
	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello with middleware")
	})

	r.Run()
}
