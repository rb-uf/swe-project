package handlers

import "github.com/gin-gonic/gin"

// This function will call the rest of the handlers

func MasterHandler(r *gin.Engine) {
	Get(r)
}

func Get(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}
