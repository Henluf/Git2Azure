package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/exp01/test", func(c *gin.Context) {
		time.Sleep(3 * time.Second)	
		if c.Query("EXP01-KEY") == "H3r34r3Dr4g0ns" {

			c.JSON(200, gin.H{
				"message": "You are the dragon master !!",
			})

		} else {

			c.JSON(403, gin.H{
				"error": "This is not the endpoint you're looking for !!! (No Dragons)",
			})

		}
	})
	_ = r.Run(":31415")
}
