package main 

import "github.com/gin-gonic/gin"

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	r.POST("/person", func(c *gin.Context) {
		var person Person 
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		c.JSON(200, person)
	})

	r.Run()
}
