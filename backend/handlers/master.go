package handlers

import "github.com/gin-gonic/gin"

// This function will call the rest of the handlers

func MasterHandler(r *gin.Engine) {
	r.GET("/", Get)
	r.POST("/", Post)
}
