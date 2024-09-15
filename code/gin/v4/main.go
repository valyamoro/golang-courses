package main 

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./uploaded"+file.Filename)
		c.String(http.StatusOK, "FIle %s uploaded!", file.Filename)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
		})
	}

	r.Run()
}
