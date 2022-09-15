package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

// Logger para registrar las salidas y entradas
// Se usa un paquete de log. Para instalar usar go get -u github.com/labstack/gommon/log
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := time.Now()
		path := c.Request.URL.Path
		verb := c.Request.Method

		// Process request
		c.Next()

		log.Infof("time: %v path: localhost:8080%s verb: %s", time, path, verb)
	}
}
