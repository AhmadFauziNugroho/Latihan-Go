package main

import "github.com/gin-gonic/gin"


type User struct {
	// Name string `json: "name"`
	// Email string `json: "email"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "pong",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"hello":name})
	})

	r.POST("/echo", func (c *gin.Context)  {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"receive":user,
		})
	})

	r.Run(":8080")
}