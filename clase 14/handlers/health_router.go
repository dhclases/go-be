package handlers

import "github.com/gin-gonic/gin"

func NewHealthRouter(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
}
